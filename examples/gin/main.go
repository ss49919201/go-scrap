package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(func(ctx *gin.Context) {
		fmt.Println("1")
		ctx.Next()
		fmt.Println("6")
	})
	engine.Use(func(ctx *gin.Context) {
		fmt.Println("2")
		ctx.Next()
		fmt.Println("5")
	})
	engine.GET("/test", func(ctx *gin.Context) {
		fmt.Println("3")
		ctx.String(200, "test")
		ctx.Next()
		fmt.Println("4")
	})
	engine.Run(":8080")
}
