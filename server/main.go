package main

import (
	"github.com/paulacgates/todo-app-golang/database"
	"github.com/paulacgates/todo-app-golang/routes"
)

func main() {
	database.InitDatabase()
	routes.HandleRoutes()
}
