package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware easy middleware
func LoggerMiddleware() gin.HandlerFunc {
	fmt.Println("Hello world!")
	return func(context *gin.Context) {
		fmt.Println("Hello LoggerMiddleware")
	}
}
