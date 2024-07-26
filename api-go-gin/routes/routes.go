package routes

import (
	"api-go-gin/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.Engine) {
	r.GET("/students", controllers.GetAllStudents)
}

func Handle404Routes(r *gin.Engine) {
	r.GET("/:allothershit", func(c *gin.Context) {
		pageName := c.Params.ByName("allothershit")
		c.JSON(404, gin.H{
			"error":   "page.not.found",
			"message": fmt.Sprintf("<html><h1>Página \"%s\" não encontrada :(</h1></html>", pageName),
		})
	})
}
