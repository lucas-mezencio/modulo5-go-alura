package main

import (
	"gin-rest-alura/database"
	"gin-rest-alura/models"
	"gin-rest-alura/routes"
)

func main() {
	database.ConnectWithDB()
	models.Alunos = []models.Aluno{
		{
			Nome: "Lucas Mezencio",
			CPF:  "12312312312",
			RG:   "mg12123123",
		},
		{
			Nome: "Segundo Nome",
			CPF:  "00011122233",
			RG:   "go12123123",
		},
	}
	routes.HandleRequests()
}
