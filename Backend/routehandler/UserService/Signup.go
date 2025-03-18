package userservice

import (
	"backend/database"
	utils "backend/utils"
	"encoding/json"

	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var signupRequest SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Not able to parse the body", http.StatusBadRequest)
		return
	}

	db := database.DB

	var userByUsername database.Users
	db.Raw("SELECT * from users where username = ?", signupRequest.Username).Scan(&userByUsername)

	if CheckIfUserExist(userByUsername) {
		http.Error(w, "User Exist with the username "+signupRequest.Username, http.StatusBadRequest)
		return
	}

	var userByEmail database.Users
	db.Raw("SELECT * from users where email = ?", signupRequest.Email).Scan(&userByEmail)

	if CheckIfUserExist(userByEmail) {
		http.Error(w, "User Exist with the email "+signupRequest.Email, http.StatusBadRequest)
		return
	}

	var user database.Users

	user.Email = signupRequest.Email
	user.Name = signupRequest.Name
	user.Username = signupRequest.Username
	user.Password = signupRequest.Password

	result := db.Create(&user)

	var signupResponse SignupResponse

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	signupResponse.Message = "User Created Successfully"
	signupResponse.UserId = user.ID

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	signupResponse.Token.UserId = user.ID
	signupResponse.Token.TokenString = token

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(signupResponse)

}

func CheckIfUserExist(user database.Users) bool {
	return user.ID != 0
}
