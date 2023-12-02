package handler

import (
	"fmt"
	"todo-cognixus/database"
	"todo-cognixus/model"

	"github.com/gofiber/fiber/v2"
)

func AddTodo(ctx *fiber.Ctx) error {
	db := database.DB
	if db == nil {
		fmt.Println("why")
	}

	var body model.Todo

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	fmt.Println(body)

	err = db.Create(&body).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to add a new todo item"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Successfully added a new todo item"})
}

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

func GetAllTodo(ctx *fiber.Ctx) error {
	var todos []model.Todo
	database.DB.Find(&todos)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": todos})
}

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
