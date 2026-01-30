package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on 8080")
	err := http.ListenAndServe(":8080" , nil)
	if err != nil {
		log.Println("Error while string the server")
	}
}