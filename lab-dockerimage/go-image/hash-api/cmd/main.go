package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

type HashRequest struct {
	Text string `json:"text"`
}

type HashResponse struct {
	MD5    string `json:"md5"`
	SHA256 string `json:"sha256"`
}

func hashHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HashRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	md5Hash := md5.Sum([]byte(req.Text))
	shaHash := sha256.Sum256([]byte(req.Text))

	response := HashResponse{
		MD5:    hex.EncodeToString(md5Hash[:]),
		SHA256: hex.EncodeToString(shaHash[:]),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/hash", hashHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
