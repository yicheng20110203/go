package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	fmt.Println("Hello world!")
	return func(context *gin.Context) {
		fmt.Println("Hello LoggerMiddleware")
	}
}
