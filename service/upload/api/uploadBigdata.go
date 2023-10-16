package api

import (
	"bufio"
	"encoding/json"
	"fileserver_enc/util"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	cmn "fileserver_enc/common"
	cfg "fileserver_enc/config"
	"fileserver_enc/mq"
	dbcli "fileserver_enc/service/dbproxy/client"
	"fileserver_enc/store/ceph"
	"fileserver_enc/store/oss"
)

// DoUploadBigDataHandler ： 处理文件上传(大文件版
func DoUploadBigDataHandler(c *gin.Context) {
	errCode := 0
	defer func() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if errCode < 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg":  "上传失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg":  "上传成功",
			})
		}
	}()

	// 1. 从form表单中获得文件内容句柄
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Failed to get form data, err:%s\n", err.Error())
		errCode = -1
		return
	}
	defer file.Close()

	//// 2. 把文件内容转为[]byte
	//buf := bytes.NewBuffer(nil)
	//if _, err := io.Copy(buf, file); err != nil {
	//	log.Printf("Failed to get file data, err:%s\n", err.Error())
	//	errCode = -2
	//	return
	//}

	// 3. 构建文件元信息(部分)
	fileMeta := dbcli.FileMeta{
		FileName: head.Filename,
		//FileSha1: util.Sha1(buf.Bytes()), //　计算文件sha1
		//FileSize: int64(len(buf.Bytes())),
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 4. 将文件写入临时存储位置
	// 4.1 创建文件
	fileMeta.Location = cfg.TempLocalRootDir + fileMeta.FileName // 临时存储地址
	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		log.Printf("Failed to create file, err:%s\n", err.Error())
		errCode = -3
		return
	}
	defer newFile.Close()

	// 4.2 写入文件
	r := bufio.NewReader(file)
	buf := make([]byte, 16777216) //16MB

	fileSize := 0
	for {
		n, err := r.Read(buf)
		//fmt.Println("读入了：", n)

		if err != nil {
			if err == io.EOF {
				// 判断文件读取结束
				break
			}
			log.Printf("打开文件失败:%#v", err.Error())
		}

		//去掉aes加密直接写入文件
		_, err = newFile.Write(buf[:n])
		if err != nil {
			fmt.Println("写入出现错误", err.Error())
		}
		fileSize += n
	}
	// 4.3 记录文件大小
	fileMeta.FileSize = int64(fileSize)
	// 4.4 记录文件哈希值
	newFile.Seek(0, 0)
	hashbuf := make([]byte, 102400)
	_, err = newFile.Read(hashbuf)
	if err != nil {
		log.Println("算哈希时读取文件100k出错！")
	}
	fileMeta.FileSha1 = util.Sha1(hashbuf)

	// 5. 同步或异步将文件转移到 Ceph/OSS
	newFile.Seek(0, 0) // 游标重新回到文件头部
	if cfg.CurrentStoreType == cmn.StoreCeph {
		// 文件写入Ceph存储
		data, _ := ioutil.ReadAll(newFile)
		cephPath := cfg.CephRootDir + fileMeta.FileSha1
		_ = ceph.PutObject("userfile", cephPath, data)
		fileMeta.Location = cephPath
	} else if cfg.CurrentStoreType == cmn.StoreOSS {
		// 文件写入OSS存储
		ossPath := cfg.OSSRootDir + fileMeta.FileName
		// 判断写入OSS为同步还是异步
		if !cfg.AsyncTransferEnable {
			// TODO: 设置oss中的文件名，方便指定文件名下载
			err = oss.Bucket().PutObject(ossPath, newFile)
			if err != nil {
				log.Println(err.Error())
				errCode = -5
				return
			}
			fileMeta.Location = ossPath
		} else {
			// 写入异步转移任务队列
			//1.构建消息体
			data := mq.TransferData{
				FileHash:      fileMeta.FileSha1,
				CurLocation:   fileMeta.Location,
				DestLocation:  ossPath,
				DestStoreType: cmn.StoreOSS,
			}
			//2.消息json化
			pubData, _ := json.Marshal(data)
			//3.Publish发布
			pubSuc := mq.Publish(
				cfg.TransExchangeName,
				cfg.TransOSSRoutingKey,
				pubData,
			)
			if !pubSuc {
				// TODO: 当前发送转移信息失败，稍后重试
				log.Println("当前发送转移信息失败，稍后重试")
			} else {
				log.Println("成功publish等待消费！")
			}

			//更新为 oss上的 location
			//fileMeta.Location = ossPath
			//不用了transfer里有做！
		}
	}

	//6.  更新文件表记录
	_, err = dbcli.OnFileUploadFinished(fileMeta)
	if err != nil {
		errCode = -6
		return
	}

	// 7. 更新用户文件表
	username := c.Request.FormValue("username")
	upRes, err := dbcli.OnUserFileUploadFinished(username, fileMeta)
	if err == nil && upRes.Suc {
		errCode = 0
	} else {
		errCode = -6
	}
}
