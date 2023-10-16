package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAccessStrucHandler : 响应访问结构获取页面
func GetAccessStruc(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/accessstruc.html")
}
