package api

import (
	"fileserver_enc/common"
	dbcli "fileserver_enc/service/dbproxy/client"
	"fileserver_enc/store/oss"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// DownloadDirURLHandler : 生成文件夹下两个文件的下载地址
func DownloadDirURLHandler(c *gin.Context) {
	filehash := c.Request.FormValue("filehash")

	// 从唯一文件表查找记录（唯一文件表似乎不用特殊标记，只用给用户文件表特殊标记，毕竟是提醒用户点击另一个下载按钮的！）
	dbResp, err := dbcli.GetFileMeta(filehash)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": common.StatusServerError,
				"msg":  "server error",
			})
		return
	}

	tblFile := dbcli.ToTableFile(dbResp.Data)

	// oss下载 文件url
	// 表中的文件名是没有改的是正常的
	// 由于 upload到 oss的名字是特殊名字，所以下载时也是用特殊名字
	special_filename := strings.Replace(tblFile.FileName.String, ".", "", -1)

	log.Println(special_filename)
	signedURLFile := oss.DownloadURL(tblFile.FileAddr.String + "/enc_" + special_filename)
	log.Println(tblFile.FileAddr.String + "/enc_" + tblFile.FileName.String)
	// oss下载 密钥url
	signedURLKey := oss.DownloadURL(tblFile.FileAddr.String + "/abe_key")
	log.Println(tblFile.FileAddr.String + "/abe_key")

	//c.Data(http.StatusOK, "application/octet-stream", []byte(signedURL))
	//c.Data(http.StatusOK, "application/json", )
	c.JSON(http.StatusOK, gin.H{
		"file": signedURLFile,
		"key":  signedURLKey,
	})

}
