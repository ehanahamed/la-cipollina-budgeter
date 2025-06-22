package main

import (
	"context"
	"fmt"
	"os"

    "github.com/joho/godotenv"
 	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultPort = "3000"

func main() {
	_ = godotenv.Load() /* "go dotenv" not "godot env" */

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		fmt.Fprint(
			os.Stderr,
			`DB_URL is not set
Copy .env.example to .env or
make sure environment variables are set elsewhere`,
		)
		os.Exit(1)
	}
	dbpool, dbpoolErr := pgxpool.New(
		context.Background(),
		dbUrl,
	)
	if dbpoolErr != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error creating database connection pool: %v\n",
			dbpoolErr,
		)
		os.Exit(1)
	}
	defer dbpool.Close()

	app := fiber.New()

	app.Get("/hi", func(c *fiber.Ctx) error {
		return c.SendString("hii world!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	appListenErr := app.Listen(":" + port)
	if appListenErr != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error starting server: %v\n",
			appListenErr,
		)
	}
}
