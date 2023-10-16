package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/welcome.html")
}
