package routes

import (
	//Middleware "backend/middlewares"
	UserService "backend/routehandler/UserService"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users/signup", UserService.Signup).Methods("POST")
	router.HandleFunc("/users/login", UserService.Login).Methods("POST")

	return router
}
