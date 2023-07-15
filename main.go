package main

import (
	"bitespeedtask/config"
	"bitespeedtask/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	config.LoadEnv()
	config.InitializeDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	routes.Routes(r)

	r.Run(":8000")
}
