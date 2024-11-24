package services

import (
	"backend-test/models"
	"backend-test/repositories"
	"backend-test/utils"
	"errors"
	"github.com/google/uuid"
	"time"
)

type AuthService struct{}

func (s *AuthService) Login(email, password string) (string, error) {
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

func (s *AuthService) MerchantLogin(email, password string) (string, error) {
	merchant, err := repositories.GetMerchantByEmail(email)
	if err != nil || merchant == nil {
		return "", errors.New("merchant not found")
	}

	if merchant.Password != password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(merchant.Id)
	if err != nil {
		return "", err
	}

	// Log login activity
	err = repositories.AddHistory(models.History{
		Id:     uuid.New().String(),
		Action: "Merchant Login",
		Detail: email + " logged in as merchant",
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
