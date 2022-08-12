package main

import (
	"log"
	"net/http"
	"server/build/service"
)

func main() {

	http.HandleFunc("/hello", service.HelloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
