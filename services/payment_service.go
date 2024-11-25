package services

type PaymentServiceInterface interface {
	ProcessPayment(fromEmail, toEmail string, amount float64) error
	MerchantToBankPayment(merchantId, bankID string, amount float64) error
}
