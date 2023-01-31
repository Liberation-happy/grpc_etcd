package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPythonInfo 获取Python服务信息
func GetPythonInfo(ginCtx *gin.Context) {
	var PythonReq service.HelloRequest
	PanicIfPythonError(ginCtx.Bind(&PythonReq))
	// 从gin.Key中取出服务实例
	PythonService := ginCtx.Keys["python_server"].(service.GreeterClient)
	PythonResp, err := PythonService.SayHello(context.Background(), &PythonReq)
	PanicIfPythonError(err)
	r := res.Response{
		Data:   PythonResp,
		Status: 200,
		Msg:    "ok",
	}
	ginCtx.JSON(http.StatusOK, r)
}
