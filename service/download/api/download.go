package api

import (
	"fileserver_enc/common"
	dbcli "fileserver_enc/service/dbproxy/client"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	// dlcfg "fileserver_enc/service/download/config"
)

// DownloadURLHandler : 生成文件的下载地址
func DownloadURLHandler(c *gin.Context) {
	filehash := c.Request.FormValue("filehash")
	// 从文件表查找记录
	//dbResp, err := dbcli.GetFileMeta(filehash)
	//if err != nil {
	//	c.JSON(
	//		http.StatusOK,
	//		gin.H{
	//			"code": common.StatusServerError,
	//			"msg":  "server error",
	//		})
	//	return
	//}

	//tblFile := dbcli.ToTableFile(dbResp.Data)

	//  判断文件存在OSS，还是在本地
	//if strings.HasPrefix(tblFile.FileAddr.String, cfg.TempLocalRootDir) {
	//	username := c.Request.FormValue("username")
	//	token := c.Request.FormValue("token")
	//	tmpURL := fmt.Sprintf("http://%s/file/download?filehash=%s&username=%s&token=%s",
	//		c.Request.Host, filehash, username, token)
	//	c.Data(http.StatusOK, "application/octet-stream", []byte(tmpURL))
	//} else if strings.HasPrefix(tblFile.FileAddr.String, cfg.OSSRootDir) {
	//	// oss下载 url (oss的 url就是提供预览 识别不了文件类型才下载就是这样,这是oss那边的问题.)
	//	// 想要实现下载得从oss get下来,然后在c.FileAttachment传递才行
	//	//signedURL := oss.DownloadURL(tblFile.FileAddr.String)
	//	//log.Println(tblFile.FileAddr.String)
	//	//c.Data(http.StatusOK, "application/octet-stream", []byte(signedURL))
	//	oss.Download(tblFile.FileAddr.String)
	//	//再像上面一样跳转本地下载 DownloadHandler
	//}

	// 不用判断在哪里了,反正本地也不删掉清理,直接走本地吧
	// 直接生成 url 对应前端的 window.open 正好
	username := c.Request.FormValue("username")
	token := c.Request.FormValue("token")
	tmpURL := fmt.Sprintf("http://%s/file/download?filehash=%s&username=%s&token=%s",
		c.Request.Host, filehash, username, token)
	c.Data(http.StatusOK, "application/octet-stream", []byte(tmpURL))
}

// DownloadHandler : 文件下载接口
func DownloadHandler(c *gin.Context) {
	fsha1 := c.Request.FormValue("filehash")
	username := c.Request.FormValue("username")
	// TODO: 处理异常情况
	fResp, ferr := dbcli.GetFileMeta(fsha1)
	ufResp, uferr := dbcli.QueryUserFileMeta(username, fsha1)
	if ferr != nil || uferr != nil || !fResp.Suc || !ufResp.Suc {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": common.StatusServerError,
				"msg":  "server error",
			})
		return
	}
	// 唯一文件表查询结果
	uniqFile := dbcli.ToTableFile(fResp.Data)
	// 用户文件表查询结果
	userFile := dbcli.ToTableUserFile(ufResp.Data)

	// 还没转移时 FileAddr中记录的是本地路径，转移后更新为 OSS路径
	// 本地文件， 直接下载
	//c.FileAttachment(uniqFile.FileAddr.String, userFile.FileName)
	// 因为 OSS 不走 url了也走从本地下载,所以这里不根据唯一文件表中记录的OSS地址了
	log.Println(uniqFile.FileAddr.String)
	log.Println("/data/fileserver/" + uniqFile.FileName.String)
	c.FileAttachment("/data/fileserver/"+uniqFile.FileName.String, userFile.FileName)

}
