package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/grpc"

// 	userPb "freedom/gen/proto/user"
// )

// type GatewayHandler struct {
// 	userClient userPb.UserServiceClient
// }

// func NewGatewayHandler(conn *grpc.ClientConn) *GatewayHandler {
// 	return &GatewayHandler{
// 		userClient: userPb.NewUserServiceClient(conn),
// 	}
// }

// // Register 处理用户注册 HTTP 请求，转发到 gRPC 服务
// func (h *GatewayHandler) Register(c *gin.Context) {
// 	var req struct {
// 		Username string `json:"username" binding:"required"`
// 		Email    string `json:"email" binding:"required,email"`
// 		Password string `json:"password" binding:"required,min=6"`
// 		Nickname string `json:"nickname"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 转换为 gRPC 请求
// 	grpcReq := &userPb.RegisterRequest{
// 		Username: req.Username,
// 		Email:    req.Email,
// 		Password: req.Password,
// 		Nickname: req.Nickname,
// 	}

// 	// 调用 gRPC 服务
// 	grpcResp, err := h.userClient.Register(c.Request.Context(), grpcReq)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 返回 HTTP 响应
// 	if grpcResp.GetBase().GetSuccess() {
// 		c.JSON(http.StatusCreated, gin.H{
// 			"success": true,
// 			"user":    grpcResp.GetUser(),
// 			"token":   grpcResp.GetToken(),
// 		})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"error":   grpcResp.GetBase().GetMessage(),
// 		})
// 	}
// }
