package handler

import (
	"todo-cognixus/database"
	"todo-cognixus/model"
)

func CreateUser(user *model.User) (bool, error) {
	var isExisting bool

	result := database.DB.
		Where(model.User{UID: user.UID}).
		FirstOrCreate(&user)

	if result.Error != nil {
		return isExisting, result.Error
	}

	if result.RowsAffected == 0 {
		isExisting = true
	}

	return isExisting, nil
}

func UpdateUser(user *model.User) error {
	return database.DB.Save(&user).Error
}
