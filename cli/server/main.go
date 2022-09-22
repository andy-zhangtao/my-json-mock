package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	_init()

	router := setRoute(gin.Default())

	router.Run(fmt.Sprintf(":%d", rp.Configure.Port))
}
