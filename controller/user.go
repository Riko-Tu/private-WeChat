package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmailLogin(c *gin.Context) {
	c.JSONP(http.StatusOK, 23)
}

func GetEmailCode(c *gin.Context) {
	email := c.Param("email")
	c.JSONP(http.StatusOK, gin.H{"email": email})
}
