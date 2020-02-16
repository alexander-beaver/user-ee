package main

import (
	"fmt"
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core/db"
	_struct "github.com/alexander-beaver/user-ee/core/struct"
	"log"
	"net/http"
)


func main() {
	sqlite := db.SetupDatabase()
	/*db.WriteErrorToDB(sqlite, _struct.Error{
		Model:        gorm.Model{},
		EndpointID:   "123",
		ErrorID:      0,
		ErrorMessage: "123",
		Comments:     "123",
		Time:         time.Now(),
	})*/

	//fmt.Println(json.Marshal(db.GetEntryFromDBGivenID(sqlite, 0)))

	_ := db.GetAllEntriesFromDB(sqlite)
	http.HandleFunc("/",api.PutErrorAPIHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
