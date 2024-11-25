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
	paymentService := &services.PaymentService{}

	// Inisialisasi controllers
	authController := &controllers.AuthController{
		AuthService: authService, // Inject AuthService ke AuthController
	}
	paymentController := &controllers.PaymentController{
		PaymentService: paymentService,
	}
	r.HandleFunc("/login", authController.LoginHandler).Methods("POST")
	r.HandleFunc("/merchant/login", authController.MerchantLoginHandler).Methods("POST")
	r.HandleFunc("/payment", paymentController.PaymentHandler).Methods("POST")
	r.HandleFunc("/merchant/payment", paymentController.MerchantPaymentHandler).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	return r
}
