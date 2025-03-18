package main

import (
	routes "backend/Routes"
	"backend/database"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	database.ConnectDB()
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	router := routes.Router()
	log.Fatal(http.ListenAndServe(":8000", corsOptions(router)))
}
