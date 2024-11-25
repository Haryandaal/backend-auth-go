package controllers

import (
	"backend-test/services"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler_Success(t *testing.T) {
	// Mock service
	mockService := &services.MockAuthService{
		LoginFunc: func(email, password string) (string, error) {
			return "mock_token", nil
		},
	}

	// Inisialisasi controller
	handler := AuthController{
		AuthService: mockService,
	}

	// Mock request and response
	body := map[string]string{
		"email":    "john@example.com",
		"password": "123456",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// run the handler
	handler.LoginHandler(rr, req)

	// Validate result
	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "mock_token", response["token"])
}

func TestMerchantLoginHandler_Success(t *testing.T) {
	// Mock service
	mockService := &services.MockAuthService{
		LoginFunc: func(email, password string) (string, error) {
			return "mock_token", nil
		},
	}

	// Inisialisasi controller
	handler := AuthController{
		AuthService: mockService,
	}

	// Mock request and response
	body := map[string]string{
		"email":    "john@example.com",
		"password": "123456",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// run the handler
	handler.MerchantLoginHandler(rr, req)

	// Validate result
	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "mock_token", response["token"])
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	// Mock service
	mockService := &services.MockAuthService{
		LoginFunc: func(email, password string) (string, error) {
			return "", errors.New("invalid credentials")
		},
	}

	handler := AuthController{
		AuthService: mockService,
	}

	body := map[string]string{
		"email":    "something@example.com",
		"password": "wrongpassword",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.LoginHandler(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid credentials")
}

func TestMerchantLoginHandler_InvalidCredentials(t *testing.T) {
	// Mock service
	mockService := &services.MockAuthService{
		LoginFunc: func(email, password string) (string, error) {
			return "", errors.New("invalid credentials")
		},
	}

	handler := AuthController{
		AuthService: mockService,
	}

	body := map[string]string{
		"email":    "something@example.com",
		"password": "wrongpassword",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.MerchantLoginHandler(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid credentials")
}

func TestLoginHandler_BadRequest(t *testing.T) {
	// Mock service
	mockService := &services.MockAuthService{}

	handler := AuthController{
		AuthService: mockService,
	}

	// Mock request with wrong body
	req := httptest.NewRequest("POST", "/login", bytes.NewReader([]byte("invalid body")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.LoginHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	//assert.Contains(t, rr.Body.String(), "Invalid request")
}

func TestMerchantLoginHandler_BadRequest(t *testing.T) {
	// Mock service
	mockService := &services.MockAuthService{}

	handler := AuthController{
		AuthService: mockService,
	}

	// Mock request with wrong body
	req := httptest.NewRequest("POST", "/login", bytes.NewReader([]byte("invalid body")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.MerchantLoginHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	//assert.Contains(t, rr.Body.String(), "Invalid request")
}
