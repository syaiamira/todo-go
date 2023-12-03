package handler

import (
	"todo-cognixus/database"
	_ "todo-cognixus/docs"
	"todo-cognixus/model"

	"github.com/gofiber/fiber/v2"
)

type IncomingTodo struct {
	Title string `json:"title"`
}

// @Summary Create new todo item
// @Description Add a new todo item
// @Tags Todo
// @Accept json
// @Produce json
// @Param IncomingTodo body IncomingTodo true "Todo object"
// @Success 200 {object} map[string]interface{}
// @Router /todo/ [post]
func AddTodo(ctx *fiber.Ctx) error {
	db := database.DB

	var body IncomingTodo

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	err = db.Create(&body).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to add a new todo item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully added a new todo item"})
}

// @Summary Get all todo items
// @Description Get all todo items
// @Tags Todo
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /todo/ [get]
func GetAllTodo(ctx *fiber.Ctx) error {
	var todos []model.Todo
	database.DB.Find(&todos)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": todos})
}

// @Summary Update a todo item status to complete
// @Description Update a todo item status to complete by ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Router /todo/complete/{id} [patch]
func UpdateTodoStatus(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = database.DB.Model(&model.Todo{}).
		Where("ID = ?", id).
		Update("IsCompleted", true).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to update todo item status"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully updated todo item status"})
}

// @Summary Delete todo item by ID
// @Description Delete a todo item by ID
// @Tags Todo
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Router /todo/{id} [delete]
func DeleteTodo(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = database.DB.Delete(&model.Todo{}, id).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to delete todo item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully deleted a todo item"})
}
