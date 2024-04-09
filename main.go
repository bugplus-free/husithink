package main

import (
	"husithink/models"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", models.SayhelloName)
	http.HandleFunc("/login", models.Login)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
