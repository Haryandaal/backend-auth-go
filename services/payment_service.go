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

func MerchantToBankPayment(merchantId, bankID string, amount float64) error {
	// Validate merchant
	merchant, err := repositories.GetMerchantById(merchantId)
	if err != nil || merchant == nil {
		return errors.New("merchant not found")
	}

	// Validate bank destination
	bank, err := repositories.GetBankByID(bankID)
	if err != nil || bank == nil {
		return errors.New("bank not found")
	}

	// Log activity
	return repositories.AddHistory(models.History{
		Id:     uuid.New().String(),
		Action: "Merchant Payment",
		Detail: merchant.Name + " transferred " + fmt.Sprintf("%.2f", amount) + " to " + bank.Name,
		Date:   time.Now().Format("2006-01-02 15:04:05"),
	})
}
