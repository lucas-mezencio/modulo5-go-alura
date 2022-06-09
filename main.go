package main

import (
	"gin-rest-alura/database"
	"gin-rest-alura/routes"
)

func main() {
	database.ConnectWithDB()
	routes.HandleRequests()
}
