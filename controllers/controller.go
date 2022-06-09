package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"nome": "Lucas Mezencio",
	})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz: ": "Ola " + nome,
	})
}
