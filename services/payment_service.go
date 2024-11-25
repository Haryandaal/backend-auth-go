package services

type PaymentServiceInterface interface {
	ProcessPayment(fromEmail, toEmail string, amount int64) error
	MerchantToBankPayment(merchantId, bankID string, amount int64) error
}
