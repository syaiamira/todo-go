package router

import (
	"todo-cognixus/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	todo := app.Group("todo")

	todo.Post("/", handler.AddTodo)
	todo.Get("/", handler.GetAllTodo)
	todo.Patch("/complete/:id", handler.UpdateTodoStatus)
	todo.Delete("/:id", handler.DeleteTodo)
}
