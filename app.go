package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexander-beaver/user-ee/api"
	"github.com/alexander-beaver/user-ee/core/db"
	"log"
	"net/http"
)


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
	res, _ := json.Marshal(db.GetAllEntriesFromDB(sqlite))
	fmt.Println(string(res))
	http.HandleFunc("/",api.PutErrorAPIHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
