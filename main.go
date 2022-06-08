package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"nome": "Lucas Mezencio",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", showStudents)
	r.Run(":5000")
}
