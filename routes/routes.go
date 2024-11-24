package routes

import (
	"backend-test/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/merchant/login", controllers.MerchantLoginHandler).Methods("POST")
	r.HandleFunc("/payment", controllers.PaymentHandler).Methods("POST")
	r.HandleFunc("/merchant/payment", controllers.MerchantPaymentHandler).Methods("POST") // Rute transaksi merchant
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	return r
}
