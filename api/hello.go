package api

import (
	"net/http"

	"myBlog/service"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Username string `form:"username"`
}

func Hello(c *gin.Context) {
	var form Message
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, service.GetMessage(form.Username))
}
