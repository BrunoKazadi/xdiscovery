package main

import (
	"log"
	"net/http"

	"github.com/BrunoKazadi/xdiscovery/api"
)

func main() {

	http.HandleFunc("/create", api.CreateRoomHandler)
	http.HandleFunc("/join", api.JoinRoomHandler)

	log.Println("Starting Server on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
