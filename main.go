package main

import (
	"husithink/models"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", models.SayhelloName)
	http.HandleFunc("/login", models.Login)
	http.HandleFunc("/submit", models.Submit)
	http.HandleFunc("/enroll", models.Enroll)
	http.HandleFunc("/upload", models.Upload)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
