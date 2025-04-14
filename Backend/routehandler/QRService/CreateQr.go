package qrservice

import (
	"backend/database"
	"backend/utils"
	"encoding/json"
	"net/http"
)

func CreateQR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Body CreateQrRequestBody
	var link database.Links
	var qr database.QR
	var links_history database.Links_history
	var response CreateQRResponse
	db := database.DB

	json.NewDecoder(r.Body).Decode(&Body)

	authHeader := r.Header.Get("Authorization")
	claims, _ := utils.VerifyJWT(authHeader)
	userId := claims["userId"].(float64)

	db.Raw(`
        INSERT INTO links (uri,user_id)
        VALUES (?, ?) returning *
    `, Body.URI, userId).Scan(&link)

	db.Raw(`
        INSERT INTO qr (name,current_link,user_id,qr_type)
        VALUES (?, ?, ?, ?) returning *
    `, Body.Name, link.ID, userId, Body.QrType).Scan(&qr)

	db.Raw(`
        INSERT INTO links_history (link_id,user_id,qr_id)
        VALUES (?, ?, ?) returning *
    `, link.ID, userId, qr.ID).Scan(&links_history)

	response.Link = link
	response.Qr = qr

	json.NewEncoder(w).Encode(response)

}

type CreateQrRequestBody struct {
	URI    string `json:"uri"`
	Name   string `json:"name"`
	QrType string `json:"qr_type"`
}

type CreateQRResponse struct {
	Qr   interface{} `json:"qr"`
	Link interface{} `json:"link"`
}
