package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"

// 	userPb "freedom/gen/proto/user"

// 	"github.com/gin-gonic/gin"
// 	"github.com/grpc-ecosystem/grpc-gateway/runtime"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// // 响应格式
// type Response struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// func main() {
// 	ctx := context.Background()

// 	// 创建网关多路复用器
// 	mux := runtime.NewServeMux()

// 	// gRPC 连接选项
// 	opts := []grpc.DialOption{
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	}

// 	// 🎉 自动注册所有 HTTP 路由！
// 	err := userPb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
// 	if err != nil {
// 		log.Fatal("注册网关失败:", err)
// 	}

// 	// 添加健康检查
// 	httpMux := http.NewServeMux()
// 	httpMux.Handle("/", mux)
// 	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(`{"status": "ok"}`))
// 	})

// 	log.Println("✅ HTTP 网关启动在 :8080")
// 	log.Println("📡 自动生成的路由:")
// 	log.Println("   POST /v1/users/register")
// 	log.Println("   POST /v1/users/login")
// 	log.Println("   GET  /v1/users/{user_id}")
// 	log.Println("   PUT  /v1/users/{user_id}")
// 	log.Println("   POST /v1/users/batch-status")

// 	if err := http.ListenAndServe(":8080", httpMux); err != nil {
// 		log.Fatal("网关启动失败:", err)
// 	}
// }

// // 健康检查
// func healthCheck(c *gin.Context) {
// 	c.JSON(http.StatusOK, Response{
// 		Code:    200,
// 		Message: "服务正常运行",
// 		Data:    gin.H{"timestamp": time.Now().Unix()},
// 	})
// }

// // func startGRPCServer(db *gorm.DB, wg *sync.WaitGroup) {
// // 	defer wg.Done()

// // 	// 创建 gRPC 服务器
// // 	grpcServer := grpc.NewServer()

// // 	// 初始化用户服务层
// // 	userRepo := userRepository.NewUserRepository(db)
// // 	userSer := userService.NewUserService(userRepo)
// // 	userHand := userHandler.NewUserHandler(userSer)

// // 	// 注册 gRPC 服务
// // 	userHandler.RegisterUserServiceServer(grpcServer, userHand)

// // 	// 启动 gRPC 服务器
// // 	lis, err := net.Listen("tcp", ":50051")
// // 	if err != nil {
// // 		log.Fatalf("failed to listen: %v", err)
// // 	}

// // 	log.Println("gRPC 服务器启动，端口号：50051")
// // 	if err := grpcServer.Serve(lis); err != nil {
// // 		log.Fatalf("failed to serve: %v", err)
// // 	}
// // }

// // func startHTTPGateway(db *gorm.DB, wg *sync.WaitGroup) {
// // 	defer wg.Done()

// // 	r := gin.Default()

// // 	// 健康检查
// // 	r.GET("/health", healthCheck)

// // 	// 创建 gRPC 客户端连接
// // 	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// // 	if err != nil {
// // 		log.Fatalf("did not connect: %v", err)
// // 	}
// // 	defer conn.Close()

// // 	// 创建网关处理器
// // 	gatewayHandler := handler.NewGatewayHandler(conn)

// // 	// 设置路由
// // 	setupRoutes(r, gatewayHandler)

// // 	log.Println("HTTP 网关启动，端口号：8080")
// // 	if err := r.Run(":8080"); err != nil {
// // 		log.Fatalf("failed to start HTTP gateway: %v", err)
// // 	}
// // }

// // func setupRoutes(r *gin.Engine, gatewayHandler *handler.GatewayHandler) {
// // 	// API v1 路由组
// // 	v1 := r.Group("/api/v1")
// // 	{
// // 		users := v1.Group("/users")
// // 		{
// // 			users.POST("/register", gatewayHandler.Register)
// // 			users.POST("/login", gatewayHandler.Login)
// // 			users.GET("/:id", gatewayHandler.GetUser)
// // 			users.PUT("/:id", gatewayHandler.UpdateUser)
// // 		}
// // 	}
// // }

// // func initDatabase() *gorm.DB {
// // 	// 数据库连接信息
// // 	username := "root"
// // 	password := "123456"
// // 	host := "localhost"
// // 	port := "3306"
// // 	dbName := "freedom"

// // 	// 构建DSN (Data Source Name)
// // 	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

// // 	// 连接数据库
// // 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// // 	if err != nil {
// // 		panic("数据库连接失败: " + err.Error())
// // 	}

// // 	return db
// // }
