package main

import (

	"log"
	"net/http"
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core"

)





func main() {
	db := core.SetupDatabase()
	http.HandleFunc("/",api.PutErrorAPIHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
