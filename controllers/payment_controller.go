package controllers

import (
	"backend-test/services"
	"backend-test/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
