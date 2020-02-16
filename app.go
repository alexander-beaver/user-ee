package main

import (
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core/db"
	"github.com/alexander-beaver/user-ee/core/struct"
	"github.com/jinzhu/gorm"
	"time"

	"log"
	"net/http"
)


func main() {
	sqlite := db.SetupDatabase()
	db.WriteErrorToDB(sqlite, _struct.Error{
		Model:        gorm.Model{},
		EndpointID:   "123",
		ErrorID:      0,
		ErrorMessage: "123",
		Comments:     "123",
		Time:         time.Now(),
	})
	http.HandleFunc("/",api.PutErrorAPIHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
