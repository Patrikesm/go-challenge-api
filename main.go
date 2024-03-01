package main

import (
	"challenge-week-one/database"
	"challenge-week-one/routes"
)

func main() {
	database.ConnectDB()

	routes.HandleRoutes()
}
