package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
)

func Hello (c *gin.Context){
	c.String(http.StatusOK, "hello gin")
}

func Test (c *gin.Context){
	// 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
	c.JSON(200, gin.H{
		"legend_data": legendData,
		"xAxis_data":  xAxisData,
	})
}