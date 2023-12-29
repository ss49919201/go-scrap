package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handleGetPing(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func handleGetPing(r gin.IRoutes) gin.IRoutes {
	r.GET("/ping", func(c *gin.Context) {
		c.Set("ping", "pong")

		c.JSON(http.StatusOK, gin.H{
			"message":                           "pong",
			`c.Get("ping")`:                     c.MustGet("ping"),
			`c.Request.Context().Value("ping")`: c.Request.Context().Value("ping"),
		})
	})
	return r
}
