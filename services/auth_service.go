package services

type AuthServiceInterface interface {
	Login(email, password string) (string, error)
	MerchantLogin(email, password string) (string, error)
}
