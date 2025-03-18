package userservice

import (
	"backend/database"
	utils "backend/utils"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var loginRequest LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)
	defer r.Body.Close()

	db := database.DB
	var user database.Users

	db.Raw("Select * from users where username = ? and password = ?", loginRequest.Username, loginRequest.Password).Scan(&user)

	if !CheckIfUserExist(user) {
		http.Error(w, "User doesnot Exist", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var loginResponse LoginResponse

	loginResponse.Message = "User " + user.Username + " is logged in successfully"
	loginResponse.Token.TokenString = token
	loginResponse.Token.UserId = user.ID

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(loginResponse)

}
