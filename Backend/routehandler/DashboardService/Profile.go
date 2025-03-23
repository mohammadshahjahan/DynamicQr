package dashboardservice

import (
	"backend/database"
	utils "backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	claims, _ := utils.VerifyJWT(authHeader)

	userId := claims["userId"]
	fmt.Println(userId)

	db := database.DB
	var user map[string]interface{}
	var user_data []map[string]interface{}
	var qr []map[string]interface{}

	db.Raw("Select * from users where id = ?", userId).Scan(&user)
	db.Raw("select *,qrs.id as qr_id, links.id as  link_id from (Select * from qr where user_id = ?) qrs join links on qrs.current_link = links.id order by updated_at desc limit 10", userId).Scan(&qr)
	db.Raw("select * from  (SELECT * FROM links JOIN links_history ON links.id = links_history.link_id WHERE links.user_id = ?) as e join qr on e.qr_id = qr.id", userId).Scan(&user_data)
	var count int32
	for _, val := range user_data {
		if c, ok := val["count"].(int32); ok {
			count += c
		} else {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	var response ProfileResponse

	response.Count = count
	response.Email = user["email"].(string)
	response.Name = user["name"].(string)
	response.Username = user["username"].(string)
	response.QRs = qr
	json.NewEncoder(w).Encode(response)

}
