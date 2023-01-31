package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetJavaInfo 获取Java服务信息
func GetJavaInfo(ginCtx *gin.Context) {
	var JavaReq service.JavaHelloRequest
	PanicIfJavaError(ginCtx.Bind(&JavaReq))
	// 从gin.Key中取出服务实例
	JavaService := ginCtx.Keys["java_server"].(service.JavaHelloServiceClient)
	JavaResp, err := JavaService.Hello(context.Background(), &JavaReq)
	PanicIfJavaError(err)
	r := res.Response{
		Data:   JavaResp,
		Status: 200,
		Msg:    "ok",
	}
	ginCtx.JSON(http.StatusOK, r)
}
