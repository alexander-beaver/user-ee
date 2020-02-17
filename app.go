package main

import (
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core"
	"github.com/alexander-beaver/user-ee/core/db"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// Returns all entries in the database
// The result will be sent back in the HTTP body
func GetAll(db2 *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		res := db.GetAllEntriesFromDB(db2)
		core.RespondOK(w, r, res)
	}

}

// Runs the server
// This is the entrypoint for the application
func main() {
	sqlite := db.SetupDatabase()
	/*
		db.WriteErrorToDB(sqlite, _struct.Error{
			EndpointID:   "123",
			ErrorID:      0,
			ErrorMessage: "123",
			Comments:     "123",
			Time:         time.Now(),
		})*/

	//fmt.Println(json.Marshal(db.GetEntryFromDBGivenID(sqlite, 0)))

	http.HandleFunc("/write", api.PutErrorAPIHandler(sqlite))
	http.HandleFunc("/db", GetAll(sqlite))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
