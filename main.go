package main

import (
	"log"
	"net/http"
	"husithink/models"
)

func main() {
	http.HandleFunc("/", models.SayhelloName)
	http.HandleFunc("/login", models.Login)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
