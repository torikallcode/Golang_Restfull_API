package handlers

import (
	"encoding/json"
	"golang_restfull_API/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users []models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// set header response sebagi json
	w.Header().Set("Content-Type", "application/json")
	// mengubah slice users ke json dan kirim sebagai response
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// set header response sebagai json
	w.Header().Set("Content-Type", "application/json")
	// ambil parameter url (/users/{id})
	params := mux.Vars(r)
	// konversi string id ke int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}
	// cari user berdasarkan id
	for _, item := range users {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// kirim user yang ditemukan jika tidak ditemukan maka kembalikan "not found"
	http.Error(w, "user not found", http.StatusNotFound)
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
