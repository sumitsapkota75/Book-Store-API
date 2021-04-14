package main

import (
	"example.com/sumit/database"
	"example.com/sumit/routes"
)

func main() {
	// Connect to database
	database.ConnectDatabase()
	// Setup Routes and initialize router
	routes.SetUpRoutes()
}
