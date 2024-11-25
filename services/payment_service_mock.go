package services

import "errors"

type MockPaymentService struct {
	ProcessPaymentFunc func(fromEmail, toEmail string, amount float64) error
}

func (m *MockPaymentService) ProcessPayment(fromEmail, toEmail string, amount float64) error {
	if m.ProcessPaymentFunc != nil {
		return m.ProcessPaymentFunc(fromEmail, toEmail, amount)
	}
	return errors.New("not implemented")
}

func (m *MockPaymentService) MerchantToBankPayment(merchantId, bankID string, amount float64) error {
	if m.ProcessPaymentFunc != nil {
		return m.ProcessPaymentFunc(merchantId, bankID, amount)
	}
	return errors.New("not implemented")
}
