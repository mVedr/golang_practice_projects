package main

import (
	"log"
	"net/http"

	"github.com/mVedr/go_jwt_take2/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
