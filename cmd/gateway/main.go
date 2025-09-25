package main

import (
	"fmt"
	userHandler "freedom/internal/user/handler"
	userRepository "freedom/internal/user/repository"
	userService "freedom/internal/user/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {

	// 初始化数据库
	db := initDatabase()

	grpcServer := grpc.NewServer()
	// 初始化各层
	userRepo := userRepository.NewUserRepository(db)
	userSer := userService.NewUserService(userRepo)
	userHand := userHandler.NewUserHandler(userSer)
	userHandler.RegisterUserService(grpcServer, userHand)
	r := gin.Default()

	r.GET("/health", healthCheck)
	fmt.Println("启动！")

	// API v1 路由组
	// v1 := r.Group("/api/v1")
	{
		// users := v1.Group("/users")
		// {
		// 	users.POST("register", userHand.Register)
		// }
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

func initDatabase() *gorm.DB {
	// 数据库连接信息
	username := "root"
	password := "123456"
	host := "localhost"
	port := "3306"
	dbName := "freedom"

	// 构建DSN (Data Source Name)
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	return db
}
