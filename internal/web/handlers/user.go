package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello (c *gin.Context){
	c.String(http.StatusOK, "hello gin")
}

func Test (c *gin.Context){
	c.String(http.StatusOK, "hello test")
}