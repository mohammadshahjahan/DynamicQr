package userservice

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	UserId  uint   `json:"userId"`
	Token   Token
}

type Token struct {
	UserId      uint   `json:"userID"`
	TokenString string `json:"token"`
}
