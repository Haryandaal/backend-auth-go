package controllers

import (
	"backend-test/services"
	"backend-test/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	// get token from header Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validate token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := tokenParts[1]
	userId, err2 := utils.ValidateToken(token)
	if err2 != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var paymentRequest struct {
		ToEmail string `json:"to_email"`
		Amount  int64  `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = services.ProcessPayment(userId, paymentRequest.ToEmail, float64(paymentRequest.Amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful"})
}

func MerchantPaymentHandler(w http.ResponseWriter, r *http.Request) {
	// get token from header Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "missing authorization token", http.StatusUnauthorized)
		return
	}

	// Validate token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		http.Error(w, "invalid authorization token", http.StatusUnauthorized)
		return
	}

	token := tokenParts[1]
	merchantID, err := utils.ValidateToken(token)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Decode request body
	var paymentRequest struct {
		BankID string  `json:"bank_id"`
		Amount float64 `json:"amount"`
	}

	err = json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// do payment
	err = services.MerchantToBankPayment(merchantID, paymentRequest.BankID, paymentRequest.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// response success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Merchant payment successful"})
}
