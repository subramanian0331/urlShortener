package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Url Shortener Service Begins")
	app := App{}
	err := godotenv.Load("urlShortener.env")
	if err != nil {
		log.Panic("could not load the env file")
	}
	app.Init()
	appPath := os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT")
	fmt.Println(appPath)
	log.Fatal(app.fiber.Listen(appPath))
}
