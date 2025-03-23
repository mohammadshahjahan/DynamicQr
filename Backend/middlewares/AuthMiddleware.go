package middlewares

import (
	utils "backend/utils"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func AuthMiddleWare(next HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		_, tokenErr := utils.VerifyJWT(authHeader)

		if tokenErr != nil {
			http.Error(w, "Token is not valid", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
