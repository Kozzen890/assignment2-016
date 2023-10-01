package main

import (
	"log"

	"github.com/Kozzen890/assignment2-016/database"
	"github.com/Kozzen890/assignment2-016/routes"
)

func main() {
	port := ":8080"
	log.Printf("Starting server on port %s\n", port)
	database.StartDB()
	routes.StartServer().Run(port)
}