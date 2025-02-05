package handlers

import (
	"encoding/json"
	"golang_restfull_API/models"
	"net/http"
)

var users []models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// set header response sebagi json
	w.Header().Set("Content-Type", "application/json")
	// mengubah slice users ke json dan kirim sebagai response
	json.NewEncoder(w).Encode(users)
}

// set header response sebagai json
// ambil parameter dari url
// konversi string id menjadi int
// cari dan update user
// hapus user lama dari slice
// ambil data baru
// pertahankan id yang sama
// tambahkan user yang sudah di update
// kirim user yang sudah diupdate

// set header response sebagai json
// ambil parameter dari url
// konversi string id menjadi int
// cari dan hapus user
// hapus user dari slice
// http.error status no content
