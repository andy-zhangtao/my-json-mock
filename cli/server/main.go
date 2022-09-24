package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	_init()

	go mock(rp)

	router := gin.Default()
	setRoute(router)

	router.Run(fmt.Sprintf(":%d", rp.Configure.Port))
}
