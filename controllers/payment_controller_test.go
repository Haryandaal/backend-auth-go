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

func TestPaymentHandler_Success(t *testing.T) {
	//
	validateTokenMock := func(token string) (string, error) {
		return "1", nil // Token valid, user ID = 1
	}
	// Mock service
	mockService := &services.MockPaymentService{
		ProcessPaymentFunc: func(fromEmail, toEmail string, amount float64) error {
			return nil
		},
	}
	// Inisialisasi controller dengan mock service
	handler := PaymentController{
		PaymentService:    mockService,
		ValidateTokenFunc: validateTokenMock,
	}
	// Mock request dan response
	body := map[string]interface{}{
		"to_email": "jane@example.com",
		"amount":   100,
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/payment", bytes.NewReader(jsonBody))
	req.Header.Set("Authorization", "Bearer valid_token")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.PaymentHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, "Payment successful", response["message"])
}

func TestPaymentHandler_MissingAuthorization(t *testing.T) {
	// Mock service
	mockService := &services.MockPaymentService{}

	// Inisialisasi controller dengan mock service
	handler := PaymentController{
		PaymentService: mockService,
	}

	// Mock request tanpa header Authorization
	body := map[string]interface{}{
		"from_email": "john@example.com",
		"to_email":   "jane@example.com",
		"amount":     100.50,
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/payment", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Jalankan handler
	handler.PaymentHandler(rr, req)

	// Validasi hasil
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Unauthorized")
}

func TestPaymentHandler_ProcessPaymentError(t *testing.T) {
	// Mock utils untuk validasi token
	validateTokenMock := func(token string) (string, error) {
		return "1", nil // Token valid, user ID = 1
	}

	// Mock service untuk pembayaran
	mockService := &services.MockPaymentService{
		ProcessPaymentFunc: func(fromUserID, toEmail string, amount float64) error {
			return errors.New("recipient not found") // Kesalahan logika pembayaran
		},
	}

	// Inisialisasi controller dengan mock service
	handler := PaymentController{
		PaymentService:    mockService,
		ValidateTokenFunc: validateTokenMock,
	}

	// Mock request dan response
	body := map[string]interface{}{
		"to_email": "unknown@example.com",
		"amount":   100,
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/payment", bytes.NewReader(jsonBody))
	req.Header.Set("Authorization", "Bearer valid_token")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Jalankan handler
	handler.PaymentHandler(rr, req)

	// Validasi hasil
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "recipient not found")
}
