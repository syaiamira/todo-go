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
// @Security BearerAuth
func AddTodo(ctx *fiber.Ctx) error {
	db := database.DB

	// Get and parse Todo from json to struct
	var body IncomingTodo

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	todo := model.Todo{
		Title:  body.Title,
		UserID: uint(ctx.Locals("user_id").(float64)),
	}

	// Insert into Todo
	err = db.Create(&todo).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to add a new todo item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully added a new todo item"})
}

// @Summary Get all todo items by user id
// @Description Get all todo items by user id from login
// @Tags Todo
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /todo/ [get]
// @Security BearerAuth
func GetAllTodoByUserID(ctx *fiber.Ctx) error {
	userID := uint(ctx.Locals("user_id").(float64))

	var todos []model.Todo
	database.DB.
		Where("user_id = ?", userID).
		Find(&todos)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": todos})
}

// @Summary Update a todo item status to complete
// @Description Update a todo item status to complete by ID
// @Tags Todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Router /todo/complete/{id} [patch]
// @Security BearerAuth
func UpdateTodoStatus(ctx *fiber.Ctx) error {
	userID := uint(ctx.Locals("user_id").(float64))

	todoID, err := ctx.ParamsInt("todo_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.
		Model(&model.Todo{}).
		Where("id = ? AND user_id = ?", todoID, userID).
		Update("IsCompleted", true)

	if result.Error != nil || result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to update todo item status"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully updated todo item status"})
}

// @Summary Delete todo item by ID
// @Description Delete a todo item by ID
// @Tags Todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} map[string]interface{}
// @Router /todo/{id} [delete]
// @Security BearerAuth
func DeleteTodo(ctx *fiber.Ctx) error {
	userID := uint(ctx.Locals("user_id").(float64))

	todoID, err := ctx.ParamsInt("todo_id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.DB.
		Where("id = ? AND user_id = ?", todoID, userID).
		Delete(&model.Todo{}, todoID)

	if result.Error != nil || result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to delete todo item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully deleted a todo item"})
}
