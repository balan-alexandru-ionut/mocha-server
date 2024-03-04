package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()

	if err := engine.Run(":8080"); err != nil {
		fmt.Println("wow")
		panic(err)
	}
}
