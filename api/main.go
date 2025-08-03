package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"la-cipollina-budgeter-api/auth"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type, Accept, Eh-Share-Link-Token",
	}))

	app.Post("/log-in", auth.LoginHandler)

	app.Get(
		"/employees",
		auth.AuthMiddleware,
		handlers.GetEmployees,
	)
	app.Post(
		"/employees",
		auth.AuthMiddleware,
		handlers.AddEmployee,
	)
	app.Put(
		"/employees/:id",
		auth.AuthMiddleware,
		handlers.UpdateEmployee,
	)
	app.Delete(
		"/employees/:id",
		auth.AuthMiddleware,
		handlers.RemoveEmployee,
	)
	/* any authed user can get all users */
	app.Get(
		"/users",
		auth.AuthMiddleware,
		handlers.GetUsers,
	)
	/* only admin users can POST or DELETE users
	(authorization logic in auth.AdminMiddleware) */
	app.Post(
		"/users",
		auth.AuthMiddleware,
		auth.AdminMiddleware,
		handlers.AddUser,
	)
	app.Delete(
		"/users/:id",
		auth.AuthMiddleware,
		auth.AdminMiddleware,
		handlers.DeleteUser,
	)
	/* authorization logic in handlers.UpdateUser
	admin users can update any user
	non-admin users can only update themselves
	and non-admin users can not update the `admin` field
	 */
	app.Patch(
		"/users/:id",
		auth.AuthMiddleware,
		/* no admin middleware,
		auth logic is in handlers.UpdateUser */
		handlers.UpdateUser,
	)

	app.Get(
		"/authed-user/",
		auth.AuthMiddleware,
		auth.GetAuthedUser,
	)
	app.Get(
		"/days/:date",
		auth.AuthOrShareLinkMiddlewareWDateParam,
		handlers.GetDayByDate,
	)
	app.Post(
		"/days",
		auth.AuthMiddleware,
		handlers.RecordDay,
	)
	app.Put(
		"/days/:id",
		auth.AuthMiddleware,
		handlers.UpdateDay,
	)
	app.Delete(
		"/days/:id",
		auth.AuthMiddleware,
		handlers.DeleteDay,
	)
	app.Get(
		"/weeks/:date",
		auth.AuthOrShareLinkMiddlewareWDateParam,
		handlers.GetWeekByDate,
	)
	app.Get(
		"/weeks",
		auth.AuthMiddleware,
		handlers.GetAllWeeks,
	)
	app.Post(
		"/weeks",
		auth.AuthMiddleware,
		handlers.AddWeek,
	)
	app.Delete(
		"/weeks/:id",
		auth.AuthMiddleware,
		handlers.DeleteWeek,
	)
	app.Get(
		"/weeks/:date/days",
		auth.AuthOrShareLinkMiddlewareWDateParam,
		handlers.GetDaysInWeekByDate,
	)

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
