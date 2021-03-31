package db

import (
	"goweb/db/models"
)

func ConvertReturnUser(user models.User, token string) models.RUser {
	return models.RUser{
		UserID:  user.UserId,
		Phone:   user.Phone,
		Balance: user.Balance,
		Token:   token,
	}
}
