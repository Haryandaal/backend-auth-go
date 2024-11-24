package services

import "errors"

type MockAuthService struct {
	LoginFunc func(email, password string) (string, error)
}

func (m *MockAuthService) MerchantLogin(email, password string) (string, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(email, password)
	}
	return "", errors.New("not implemented")
}

func (m *MockAuthService) Login(email, password string) (string, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(email, password)
	}
	return "", errors.New("not implemented")
}
