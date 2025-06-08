package qrservice

import (
	"backend/database"
	utils "backend/utils"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

func GetQrById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	claims, _ := utils.VerifyJWT(authHeader)

	var history []map[string]interface{}
	var qrDetails map[string]interface{}
	var hasNext bool
	userId := claims["userId"].(float64)
	queryParams := r.URL.Query()
	qrID := queryParams.Get("qrID")
	offset := queryParams.Get("offset")
	db := database.DB

	if qrID == "" {
		http.Error(w, "qrID is required", http.StatusBadRequest)
		return
	}

	if offset == "" {
		http.Error(w, "offset is required", http.StatusBadRequest)
		return
	}

	offsetUint, err := strconv.ParseUint(offset, 10, 32)

	if err != nil {
		http.Error(w, "Invalid offset", http.StatusBadRequest)
		return
	}

	err = db.Raw("select * from (select * from links_history where user_id = ? and qr_id = ? ) as e join links on links.id = e.link_id limit 10 offset ?", uint(userId), qrID, offsetUint).Scan(&history).Error
	if err != nil {
		http.Error(w, "Failed to fetch QR data", http.StatusInternalServerError)
		return
	}

	err = db.Raw("select * from (select * from qr where id = ? ) as e join links on links.id = e.current_link", qrID).Scan(&qrDetails).Error
	if err != nil {
		http.Error(w, "Failed to fetch QR data", http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(len(history)) / float64(10)))
	newOffset := offsetUint + 10

	if newOffset < uint64(len(history)) {
		hasNext = true
	} else {
		hasNext = false
	}
	var reponse QRHistoryResponse

	reponse.Details = qrDetails
	reponse.History = history
	reponse.Info.HasNext = hasNext
	reponse.Info.NewOffset = int(newOffset)
	reponse.Info.TotalPages = totalPages

	json.NewEncoder(w).Encode(reponse)
}

type QRHistoryResponse struct {
	Details map[string]interface{}   `json:"details"`
	History []map[string]interface{} `json:"history"`
	Info    Info                     `json:"info"`
}
