package main

import (
	"log"
	"os"

    "github.com/joho/godotenv"
 	"github.com/gofiber/fiber/v2"
	"la-cipollina-budgeter-api/db"
)

const defaultPort = "3000"

func main() {
	_ = godotenv.Load() /* "go dotenv" not "godot env" */
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db.Connect()
	defer db.Pool.Close()

	app := fiber.New()

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.SendString("hii world!")
	})

	appListenErr := app.Listen(":"+port)
	if appListenErr != nil {
		log.Println("Error starting server: ", appListenErr)
	}
}
