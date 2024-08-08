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
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao processar os dados",
		})
	}

	student.ID = ulid.Make().String()
	models.Alunos = append(models.Alunos, student)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso!",
	})
}
