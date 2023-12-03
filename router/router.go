package router

import (
	"todo-cognixus/handler"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	SetupSwagger(app)

	todo := app.Group("todo")

	todo.Post("/", handler.AddTodo)
	todo.Get("/", handler.GetAllTodo)
	todo.Patch("/complete/:id", handler.UpdateTodoStatus)
	todo.Delete("/:id", handler.DeleteTodo)
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://127.0.0.1:8000/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))
}
