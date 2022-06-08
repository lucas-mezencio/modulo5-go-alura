package routes

import (
	"gin-rest-alura/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllStudents)
	_ = r.Run(":5000")
}
