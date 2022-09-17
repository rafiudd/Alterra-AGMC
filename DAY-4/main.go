package main

import (
	"day_4/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":8080"
	e := routes.NewRoutes()
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Printf("Successfully started on port %s\n", port)
}
