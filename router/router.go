package router

import (
	"todo-cognixus/handler"
	"todo-cognixus/middleware"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	SetupSwagger(app)

	// Auth group routers
	auth := app.Group("auth")

	auth.Get("/:provider", handler.Login)
	auth.Get("/:provider/callback", handler.LoginCallback)

	// Todo group routers
	todo := app.Group("todo")

	todo.Post("/", middleware.ValidateToken(), handler.AddTodo)
	todo.Get("/", middleware.ValidateToken(), handler.GetAllTodoByUserID)
	todo.Patch("/complete/:todo_id", middleware.ValidateToken(), handler.UpdateTodoStatus)
	todo.Delete("/:todo_id", middleware.ValidateToken(), handler.DeleteTodo)
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
