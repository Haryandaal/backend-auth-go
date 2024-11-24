package services

import (
	"backend-test/models"
	"backend-test/repositories"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func ProcessPayment(fromEmail, toEmail string, amount float64) error {
	// validate user/recipient
	toCustomer, err := repositories.GetCustomerByEmail(toEmail)
	if err != nil || toCustomer == nil {
		return errors.New("recipient not found")
	}

	// Log payment activity
	return repositories.AddHistory(models.History{
		Id:     uuid.New().String(),
		Action: "Payment",
		Detail: fromEmail + " sent " + toEmail + "an amount of " + fmt.Sprintf("%.2f", amount),
		Date:   time.Now().Format("2006-01-02 15:04:05"),
	})

}
