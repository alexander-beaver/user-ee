package main

import (
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core/db"
	_struct "github.com/alexander-beaver/user-ee/core/struct"
	"log"
	"net/http"
	"time"
)


func main() {
	sqlite := db.SetupDatabase()
	db.WriteErrorToDB(sqlite, _struct.Error{
		EndpointID:   "123",
		ErrorID:      0,
		ErrorMessage: "123",
		Comments:     "123",
		Time:         time.Now(),
	})

	//fmt.Println(json.Marshal(db.GetEntryFromDBGivenID(sqlite, 0)))

	http.HandleFunc("/",api.PutErrorAPIHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
