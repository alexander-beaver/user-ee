package main

import (
	"log"
	"net/http"
	"github.com/alexander-beaver/user-ee/api"

)





func main() {
	http.HandleFunc("/", )

	log.Fatal(http.ListenAndServe(":8080", nil))
}
