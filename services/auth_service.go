package services

import (
	"backend-test/models"
	"backend-test/repositories"
	"backend-test/utils"
	"errors"
	"github.com/google/uuid"
	"time"
)

func Login(email, password string) (string, error) {
	customer, err := repositories.GetCustomerByEmail(email)
	if err != nil || customer == nil {
		return "", errors.New("customer not found")
	}

	if customer.Password != password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(customer.Id)
	if err != nil {
		return "", err
	}

	// Log login activity
	err = repositories.AddHistory(models.History{
		Id:     uuid.New().String(),
		Action: "Login",
		Detail: email + " logged in",
		Date:   time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func Logout(email string) error {
	return repositories.AddHistory(models.History{
		Id:     uuid.New().String(),
		Action: "Logout",
		Detail: email + " logged out",
		Date:   time.Now().Format("2006-01-02 15:04:05"),
	})
}
