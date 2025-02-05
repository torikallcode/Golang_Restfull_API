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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// set header response sebagai json
	w.Header().Set("Content-Type", "application/json")
	// siapkan variable untuk menampung input
	var user models.User
	// decode body request json ke struct user
	json.NewDecoder(r.Body).Decode(&user)
	// generate ID sederhana dari slice users
	user.ID = len(users) + 1
	// tambahkan user yang baru dibuat
	users = append(users, user)
	// kirim user yang sudah ditambahkan
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// set header response sebagai json
	w.Header().Set("Content-Type", "application/json")
	// ambil parameter dari url
	params := mux.Vars(r)
	// konversi string id menjadi int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
	// cari dan update user
	for index, item := range users {
		if item.ID == id {
			users = append(users[:index], users[index+1:]...)
			var user models.User
			_ = json.NewDecoder(r.Body).Decode(&user)
			item.ID = id
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

// set header response sebagai json
// ambil parameter dari url
// konversi string id menjadi int
// cari dan hapus user
// hapus user dari slice
// http.error status no content
