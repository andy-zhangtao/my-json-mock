package main

import (
	"fmt"
	"github.com/andy-zhangtao/my-json-mock/endpoints"
	"github.com/andy-zhangtao/my-json-mock/transports"
	"github.com/andy-zhangtao/my-json-mock/types"
	"github.com/andy-zhangtao/my-json-mock/types/config"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/transport/http"
)

// mock.go 是给用户使用的服务
// 用户提交mock请求后，mock.go
// 会接管所有来自用户的请求，然后
// 返回指定的内容

func mock(rp config.RunParams) {
	router := gin.Default()
	mockRouter(router)

	router.Run(fmt.Sprintf(":%d", rp.Configure.MockPort))
}

// mockRouter 拦截所有的请求
func mockRouter(router *gin.Engine) {
	runMock := endpoints.MakeRunMockEndpoint(IMockServices())

	runMockHandler := http.NewServer(runMock, transports.RunMockDecodeRequest, transports.RunMockEncodeResponse)

	router.GET(types.RunMock, gin.WrapH(runMockHandler))

	//func(context *gin.Context) {
	//	logrus.Debugf("%s", context.FullPath())
	//	logrus.Debugf("%s", context.Param("url"))
	//	logrus.Debugf("%s", context.Param("id"))
	//	context.JSON(200, "{\"code\":300}")
	//}
	return
}
