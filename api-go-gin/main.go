package main

import (
	"api-go-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine = gin.Default()

	routes.HandleRoutes(r)
	routes.Handle404Routes(r)

	r.Run(":8080")
}
