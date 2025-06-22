package main

import (
	"log"
	"os"

    "github.com/joho/godotenv"
 	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/handlers"
)

const defaultPort = "3000"

func main() {
	_ = godotenv.Load() /* "go dotenv" not "godot env" */

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal(
			`DB_URL is not set
Copy .env.example to .env or
just check your environment variables`,
		)
	}
	db.Init(dbUrl)
	defer db.Pool.Close()

	app := fiber.New()

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.SendString("hii world!")
	})
	app.Get("/employees", handlers.GetEmployees)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(
			"Error starting server: ",
			err,
		)
	}
}
