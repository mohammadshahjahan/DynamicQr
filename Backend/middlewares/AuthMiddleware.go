package middlewares

import (
	utils "backend/utils"
	"encoding/json"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func AuthMiddleWare(next HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		token, ok := body["token"].(string)

		if !ok || token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		_, tokenErr := utils.VerifyJWT(token)

		if tokenErr != nil {
			http.Error(w, "Token is not valid", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
