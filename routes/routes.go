package routes

import (
	"gin-rest-alura/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllStudents)
	r.GET("/:nome", controllers.Greeting)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.GET("/alunos/:id", controllers.FindStudentById)
	r.DELETE("/alunos/:id", controllers.DeleteStudentById)
	_ = r.Run(":5000")
}
