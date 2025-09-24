package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()

	r.GET("/health", healthCheck)
	fmt.Println("启动！")

	// API v1 路由组
	// v1 := r.Group("/api/v1")
	{

	}

	log.Println("Freedom 服务启动，端口号：8080")
	if err := r.Run(":8080"); err != nil {
		log.Println("服务启动失败")
	}
}

// 健康检查
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "服务正常运行",
		Data:    gin.H{"timestamp": time.Now().Unix()},
	})
}
