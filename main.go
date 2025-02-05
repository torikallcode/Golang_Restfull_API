package main

import (
	"golang_restfull_API/router"
	"log"
	"net/http"
)

func main() {
	router := router.SetupRouter()

	log.Println("Server sedang berjalan di port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
