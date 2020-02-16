package main

import (
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core/db"
	"log"
	"net/http"
)





func main() {
	db := db.SetupDatabase()
	http.HandleFunc("/",api.PutErrorAPIHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
