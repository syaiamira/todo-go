package main

import (
	"log"
	"todo-cognixus/database"
	_ "todo-cognixus/docs"
	"todo-cognixus/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Todo
// @version 1.0
// @description This is a swagger for Todo
// @termsOfService http://swagger.io/terms/
// @contact.name Amira
// @contact.email syahidatulamira06@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8000
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	database.ConnectSQLite()
	router.SetupRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
