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
