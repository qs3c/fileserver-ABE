package api

import (
	cfg "fileserver_enc/config"
	"fileserver_enc/service/abe/config"
	"fileserver_enc/service/abe/test"
	dbcli "fileserver_enc/service/dbproxy/client"
	"fileserver_enc/store/oss"
	multiencwrite "fileserver_enc/testMultiEnc/api"
	"fileserver_enc/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// 想要使用服务，首先需要客户端
//var (
//	abeCli abeProto.EncUploadService
//)
//
//func init() {
//	//配置请求容量及qps
//	bRate := ratelimit2.NewBucketWithRate(100, 1000)
//	myRegistry := consul.NewRegistry(registry.Addrs("localhost:8500"))
//	service := micro.NewService(
//		micro.Flags(cmn.CustomFlags...),
//		micro.WrapClient(ratelimit.NewClientWrapper(bRate, false)), //加入限流功能, false为不等待(超限即返回请求失败)
//		micro.WrapClient(hystrix.NewClientWrapper()),               // 加入熔断功能, 处理rpc调用失败的情况(cirucuit breaker)
//		micro.Registry(myRegistry),
//	)
//	// 初始化， 解析命令行参数等
//	service.Init()
//
//	// 初始化一个abe加密服务的客户端
//	abeCli = abeProto.NewEncUploadService("go.micro.service.abe", service.Client())
//
//}

// // EncUploadHandler ： 处理文件加密上传
func EncUploadBigDataHandler(c *gin.Context) {
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

	//0. 从head中获取accessPolicy 和 username
	accessPolicy := c.Request.FormValue("accessPolicy")
	username := c.Request.FormValue("username")

	// 1. 从form表单中获得文件内容句柄
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Failed to get form data, err:%s\n", err.Error())
		errCode = -1
		return
	}
	defer file.Close()

	// 2. 把文件内容转为[]byte → dataBlocks
	//r := bufio.NewReader(file)
	////var dataBlocks []byte
	//dataBlocks := make([]byte, 52428800) //50MB
	//buf := make([]byte, 102400)          // 每次读取字节数 100KB
	//for {
	//	n, err := r.Read(buf) // 读取字节数 n
	//	if err != nil {
	//		if err == io.EOF {
	//			// 判断文件读取结束
	//			break
	//		}
	//		log.Printf("打开文件失败:%#v", err.Error())
	//	}
	//	dataBlocks = append(dataBlocks, buf[:n]...) // 注意有人这里[:n] 是读的字节数赋值，最后一次读取可能小于buf定义量
	//}

	// 3. 构建文件元信息(部分，没有哈希值和文件大小)
	fileMeta := dbcli.FileMeta{
		FileName: head.Filename,
		//FileSha1: util.Sha1(dataBlocks), //　计算文件sha1
		//FileSize: int64(len(dataBlocks)),
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 4. 在目标路径下创建文件夹：username + "_" + head.Filename
	destDirName := cfg.TempLocalRootDir + username + "_" + head.Filename
	//tempDestDirNameForLinux := "/root/ABE-master/test/temp"
	err = os.Mkdir(destDirName, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	//5. 调用ABE服务获取对称加密 key (rpc服务) 获取到对称加密 key
	//rpcData := accessPolicy + "#" + destDirName
	//log.Println("运行到这里了！")
	//resp, err := abeCli.AbEncryption(context.TODO(), &abeProto.ReqAccessPolicy{
	//	AccessPolicy: rpcData,
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//	c.Status(http.StatusInternalServerError)
	//	return
	//}
	//log.Println(resp.Key)

	// 世上无难事只要肯放弃
	// 5. 获取key -非服务直接调用函数版
	rpcData := accessPolicy + "#" + destDirName
	test.AbEncryption(rpcData)
	key := []byte(config.Key)

	// 5.1 key 取前 16字节，也就是256bit
	key_128bit := key[:16]
	//key_256bit := key[:32]
	log.Println("这是256b的key", key_128bit)
	log.Println("下面要开始aes加密了")
	// 6. aes加密
	//enc_file_stream, err := aesutil.EncryptAES(dataBlocks, key_128bit)
	//log.Println(len(dataBlocks))
	//enc_file_stream, err := aesutil.EncryptAES(dataBlocks, key_128bit)
	//log.Println(len(enc_file_stream))
	//if err != nil {
	//	log.Println("aes加密时发生错误！", err.Error())
	//}
	//log.Println("aes加密已完成")
	// 6.1 将加密后的文件哈希值写入元信息
	//fileMeta.FileSha1 = util.Sha1(enc_file_stream)
	//log.Println(fileMeta.FileSha1)

	// 7. 将加密文件写入目标目录（唯一文件表中记录目标目录位置）
	//fileMeta.Location = destDirName + "/enc_" + fileMeta.FileName // 临时存储地址
	// 7.1唯一文件表中记录目标目录位置
	//log.Println("准备将加密文件写入")
	//fileMeta.Location = destDirName
	//encFile, err := os.Create(destDirName + "/enc_" + fileMeta.FileName)
	//if err != nil {
	//	log.Printf("Failed to create file, err:%s\n", err.Error())
	//	errCode = -3
	//	return
	//}
	//defer encFile.Close()
	//
	//nByte, err := encFile.Write(enc_file_stream)
	//if err != nil {
	//	log.Printf("Failed to save data into file, writtenSize:%d, err:%s\n", nByte, err.Error())
	//	errCode = -4
	//	return
	//}

	//不用aes工具里的写入了，md写不进去！
	//aesutil.WriteStringToFile(destDirName+"/enc_"+fileMeta.FileName, string(enc_file_stream))
	// 从file文件句柄中循环读取数据，做aes加密，然后写入目标路径中的文件里，并返回文件大小
	fileMeta.FileSize, err = multiencwrite.MultiEncWrite(key_128bit, file, destDirName+"/enc_"+fileMeta.FileName)
	if err != nil {
		log.Println("写入aes加密文件失败")
		errCode = -2
		return

	}

	// 8. 直接转移至OSS
	// 8.1 打开加密文件和 key
	fileMeta.Location = destDirName
	newFile, err := os.Open(fileMeta.Location + "/enc_" + fileMeta.FileName)
	if err != nil {
		log.Printf("Failed to open file, err:%s\n", err.Error())
		errCode = -3
		return
	}
	defer newFile.Close()

	//filehash 在这里完成,取头部 100kb进行哈希
	newFile.Seek(0, 0)
	hashbuf := make([]byte, 102400)
	newFile.Read(hashbuf)
	fileMeta.FileSha1 = util.Sha1(hashbuf)

	//newKey, err := os.Open(destDirName+"/abe_key")
	newKey, err := os.Open(fileMeta.Location + "/abe_key")
	if err != nil {
		log.Printf("Failed to open file, err:%s\n", err.Error())
		errCode = -4
		return
	}
	defer newKey.Close()

	// 8.2 转移 OSS （不需要创建文件夹）
	newFile.Seek(0, 0) // 游标重新回到文件头部
	newKey.Seek(0, 0)  // 游标重新回到文件头部

	// username+"_"+fileMeta.FileNam 是文件夹名
	// ”enc_“+fileMeta.FileName 是文件名
	// OSS的目标文件夹 路径

	ossDestDirPath := cfg.OSSRootDir + username + "_" + fileMeta.FileName

	// 修改元信息中记录的目标文件夹地址为 oss上的文件夹
	fileMeta.Location = ossDestDirPath

	// 8.2.1.转移加密文件
	// 特殊处理文件名：去掉点.
	//Trim只能去首尾的还得是Repalce,-1代表换掉所有的不指定具体个数
	special_filename := strings.Replace(fileMeta.FileName, ".", "", -1)
	log.Println(special_filename)
	err = oss.Bucket().PutObject(ossDestDirPath+"/enc_"+special_filename, newFile)
	if err != nil {
		log.Println(err.Error())
		errCode = -5
		return
	}

	// 8.2.2.转移ABE加密 key
	err = oss.Bucket().PutObject(ossDestDirPath+"/abe_key", newKey)
	if err != nil {
		log.Println(err.Error())
		errCode = -6
		return
	}

	// 9. 更新两表
	// 9.1 更新唯一文件表
	_, err = dbcli.OnFileUploadFinished(fileMeta)
	if err != nil {
		errCode = -7
		return
	}

	// 9.2 更新用户文件表
	//记录到用户文件表时需要特殊标记 +"[enc]"
	fileMeta.FileName = "[enc]" + fileMeta.FileName
	upRes, err := dbcli.OnUserFileUploadFinished(username, fileMeta)
	if err == nil && upRes.Suc {
		errCode = 0
	} else {
		errCode = -8
	}

}
