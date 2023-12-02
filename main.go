package main

import (
	"log"
	"todo-cognixus/database"
	"todo-cognixus/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	database.ConnectSQLite()
	router.SetupRoutes(app)

	err := app.Listen("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
}
