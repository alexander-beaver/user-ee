package main

import (
	"log"
	"net/http"
	"./contexts/api"
)

func main() {
	http.HandleFunc("/", )
	log.Fatal(http.ListenAndServe(":8080", nil))
}
