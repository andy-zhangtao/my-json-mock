package main

import (
	"github.com/andy-zhangtao/my-json-mock/services"
	"github.com/andy-zhangtao/my-json-mock/types"
	"github.com/andy-zhangtao/my-json-mock/types/config"
	"log"
	"os"
)

var rp config.RunParams
var ms services.MockService

func _init() {
	if os.Getenv(types.MyJsonMockConfig) == "" {
		log.Panicf("%s empty", types.MyJsonMockConfig)
		os.Exit(-1)
	}

	data, err := os.ReadFile(os.Getenv(types.MyJsonMockConfig))
	if err != nil {
		log.Panicf("%s", err.Error())
		os.Exit(-1)
	}

	rp, err = config.ParseParams(data)
	if err != nil {
		log.Panicf("%s", err.Error())
		os.Exit(-1)
	}

	config.Output(rp)

	ms, err = services.NewMockService(rp)
	if err != nil {
		log.Panicf(err.Error())
	}
}
