package qrservice

import (
	"backend/database"
	"backend/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func UpdateQrByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body UpdateQrRequestBody
	var link database.Links
	var updatedQR database.QR
	var links_history database.Links_history
	var response UpdateQRResponse
	db := database.DB

	json.NewDecoder(r.Body).Decode(&body)

	authHeader := r.Header.Get("Authorization")
	claims, _ := utils.VerifyJWT(authHeader)
	userId := claims["userId"].(float64)

	idInt, err := strconv.Atoi(body.Id)
	if err != nil {
		http.Error(w, "qr_id must be an integer", http.StatusBadRequest)
		return
	}

	db.Raw(`
        INSERT INTO links (uri,user_id)
        VALUES (?, ?) returning *
    `, body.Uri, userId).Scan(&link)

	db.Raw(`
		UPDATE qr 
		SET current_link = ?, updated_at = ? 
		WHERE user_id = ? AND id = ? 
		RETURNING *`,
		link.ID, time.Now(), userId, idInt,
	).Scan(&updatedQR)

	db.Raw(`
        INSERT INTO links_history (link_id,user_id,qr_id)
        VALUES (?, ?, ?) returning *
    `, link.ID, userId, updatedQR.ID).Scan(&links_history)

	response.Link = link
	response.Qr = updatedQR

	json.NewEncoder(w).Encode(response)
}

type UpdateQrRequestBody struct {
	Id  string `json:"qr_id"`
	Uri string `json:"uri"`
}

type UpdateQRResponse struct {
	Qr   interface{} `json:"qr"`
	Link interface{} `json:"link"`
}
