package proxyurl

import (
	"backend/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ProxyUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	qrId, _ := strconv.ParseUint(params["qrID"], 10, 32)
	db := database.DB

	var qr database.QR
	var link database.Links
	db.Raw("SELECT * from qr where id = ?", qrId).Scan(&qr)
	db.Raw("SELECT * from links where id = ?", qr.CurrentLink).Scan(&link)
	db.Exec("UPDATE links SET count = ? WHERE id = ?", link.Count+1, link.ID)

	type responseStruct struct {
		Curr_link string `json:"uri"`
	}

	var resposne responseStruct
	resposne.Curr_link = link.Uri

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resposne)
}

// login/signup  -> user profile -> all qrs mini version -> dedicated qrs -> single qr
// edit qr -> delete qr -> create qr
