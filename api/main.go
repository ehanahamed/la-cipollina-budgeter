package main

import (
	"log"
	"os"

    "github.com/joho/godotenv"
 	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/auth"
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

	app.Use(cors.New(cors.Config{
	    AllowOrigins: "*",
	    AllowHeaders: "Authorization, Content-Type, Accept",
	}))

	app.Post("/log-in", auth.LoginHandler)

	app.Use(auth.AuthMiddleware) /* routes before this don't need auth,
	routes after this only run the handlers if the user is logged in */

	app.Get("/employees", handlers.GetEmployees)
	app.Post("/employees", handlers.AddEmployee)
	app.Put("/employees/:id", handlers.UpdateEmployee)
	app.Delete("/employees/:id", handlers.RemoveEmployee)
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.AddUser)
	app.Patch("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)
	app.Get("/authed-user/", auth.GetAuthedUser)

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
