package main

import (
	"github.com/andy-zhangtao/my-json-mock/endpoints"
	"github.com/andy-zhangtao/my-json-mock/transports"
	"github.com/andy-zhangtao/my-json-mock/types"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/transport/http"
)

func setRoute(router *gin.Engine) *gin.Engine {
	router = mockRoute(router)
	return router
}

func mockRoute(router *gin.Engine) *gin.Engine {
	findAllMock := endpoints.MakeFindAllMockEndpoint(IMockServices())
	findSpecifyMockWithId := endpoints.MakeFindSpecifyMockWithIdEndpoint(IMockServices())
	addNewMock := endpoints.MakeAddNewMockEndpoint(IMockServices())
	updateMock := endpoints.MakeUpdateMockEndpoint(IMockServices())
	deleteMock := endpoints.MakeDeleteMockEndpoint(IMockServices())

	findAllMockHandler := http.NewServer(findAllMock, transports.FindAllMockDecodeRequest, transports.FindAllMockEncodeResponse)
	findSpecifyMockWithIdHandler := http.NewServer(findSpecifyMockWithId, transports.FindSpecifyMockWithIdDecodeRequest, transports.FindSpecifyMockWithIdEncodeResponse)
	addNewMockHandler := http.NewServer(addNewMock, transports.AddNewMockDecodeRequest, transports.AddNewMockEncodeResponse)
	updateMockHandler := http.NewServer(updateMock, transports.UpdateMockDecodeRequest, transports.UpdateMockEncodeResponse)
	deleteMockHandler := http.NewServer(deleteMock, transports.DeleteMockDecodeRequest, transports.DeleteMockEncodeResponse)

	router.GET(types.GetAllMock, gin.WrapH(findAllMockHandler))
	router.GET(types.GetSpecMock, gin.WrapH(findSpecifyMockWithIdHandler))
	router.POST(types.AddNewMock, gin.WrapH(addNewMockHandler))
	router.POST(types.UpdateMock, gin.WrapH(updateMockHandler))
	router.DELETE(types.DeleMock, gin.WrapH(deleteMockHandler))
	return router
}
