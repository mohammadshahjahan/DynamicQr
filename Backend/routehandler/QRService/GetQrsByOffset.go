package qrservice

import (
	"backend/database"
	utils "backend/utils"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func GetQrsByOffest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	claims, _ := utils.VerifyJWT(authHeader)

	userId := claims["userId"].(float64)
	queryParams := r.URL.Query()
	offset := queryParams.Get("offset")
	qr_type := queryParams.Get("type")

	if offset == "" {
		http.Error(w, "offset is required", http.StatusBadRequest)
		return
	}
	offsetUint, err := strconv.ParseUint(offset, 10, 32)
	if err != nil {
		http.Error(w, "Invalid offset", http.StatusBadRequest)
		return
	}

	var reponse Response

	var qrs []map[string]interface{}

	db := database.DB

	err = db.Raw("select *,qrs.id as qr_id, links.id as  link_id from (Select * from qr where user_id = ? and qr_type = ?) qrs join links on qrs.current_link = links.id order by updated_at desc limit 10 offset ?", uint(userId), qr_type, offsetUint).Scan(&qrs).Error
	if err != nil {
		http.Error(w, "Failed to fetch QR data", http.StatusInternalServerError)
		return
	}
	reponse.Qrs = qrs

	var totalRows int
	err = db.Raw("SELECT COUNT(*) FROM qr WHERE user_id = ? AND qr_type = ?", uint(userId), qr_type).Scan(&totalRows).Error
	if err != nil {
		http.Error(w, "Failed to count QR data", http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(10)))
	fmt.Println("Total Pages: ", totalPages)
	newOffset := offsetUint + 10
	var hasNext bool
	if newOffset < uint64(totalRows) {
		hasNext = true
	} else {
		hasNext = false
	}
	reponse.Info.HasNext = hasNext
	reponse.Info.NewOffset = int(newOffset)
	reponse.Info.TotalPages = int(totalPages)

	json.NewEncoder(w).Encode(reponse)

}

type Response struct {
	Qrs  []map[string]interface{} `json:"qrs"`
	Info Info                     `json:"info"`
}

type Info struct {
	HasNext    bool `json:"has_next"`
	TotalPages int  `json:"total_pages"`
	NewOffset  int  `json:"offset"`
}
