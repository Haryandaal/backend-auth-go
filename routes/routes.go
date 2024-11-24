package routes

import (
	"backend-test/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/payment", controllers.PaymentHandler).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	return r
}
