package userservice

import (
	"encoding/json"
	"net/http"
)

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response CheckAuthResponse
	response.Status = "success"
	response.Message = "User is authenticated"
	json.NewEncoder(w).Encode(response)
}

type CheckAuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
