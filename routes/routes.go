package routes

import (
	"backend-test/controllers"
	"backend-test/services"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Inisialisasi services
	authService := &services.AuthService{} // Gunakan implementasi asli AuthService

	// Inisialisasi controllers
	authController := &controllers.AuthController{
		AuthService: authService, // Inject AuthService ke AuthController
	}
	r.HandleFunc("/login", authController.LoginHandler).Methods("POST")
	r.HandleFunc("/merchant/login", authController.MerchantLoginHandler).Methods("POST")
	r.HandleFunc("/payment", controllers.PaymentHandler).Methods("POST")
	r.HandleFunc("/merchant/payment", controllers.MerchantPaymentHandler).Methods("POST") // Rute transaksi merchant
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	return r
}
