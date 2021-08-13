package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zazin/go-country/src/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"success": true,
			"version": "1.0.0",
		})
	})
	country := controllers.New()
	r.GET("/country", country.GetUsers)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
