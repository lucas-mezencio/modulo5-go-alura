package controllers

import (
	"gin-rest-alura/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz: ": "Ola " + nome,
	})
}
