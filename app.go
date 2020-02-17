package main

import (
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core"
	"github.com/alexander-beaver/user-ee/core/db"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func WriteJSONDB(db2 *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := db.GetAllEntriesFromDB(db2)
		core.RespondOK(w, r, res)
	}

}

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

	http.HandleFunc("/", api.PutErrorAPIHandler)
	http.HandleFunc("/db", WriteJSONDB(sqlite))
	http.HandleFunc("/write", WriteJSONDB(sqlite))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
