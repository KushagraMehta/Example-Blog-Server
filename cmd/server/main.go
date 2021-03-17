package main

import (
	"fmt"
	"log"
	"os"

	controllers "github.com/KushagraMehta/Example-Blog-Server/pkg/controller"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	server.Initialize()

	server.Run(":" + os.Getenv("SERVER_PORT"))
	defer server.DB.Close()
}
