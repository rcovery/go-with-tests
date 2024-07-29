package controllers

import (
	"api-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func CreateStudent(c *gin.Context) {
	var aluno models.Student
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	aluno.ID = ulid.Make().String()

	models.Alunos = append(models.Alunos, aluno)
	c.JSON(http.StatusCreated, gin.H{
		"message": "user.created",
	})
}
