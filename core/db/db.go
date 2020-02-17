package db

import (
	"encoding/json"
	"fmt"
	"github.com/alexander-beaver/user-ee/core/struct"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// The file that will be used for the sqlite3 database
var dbName = "test.db"

// Initializes the GORM database
func SetupDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&_struct.Error{})
	db.Set("gorm:auto_preload", true)

	return db
}

// Write a new entry to the error database
func WriteErrorToDB(db *gorm.DB, reported _struct.Error) {

	db.Create(&reported)
	fmt.Println("Created Object")

}

// Gets an entry from the database given an ID
func GetEntryFromDBGivenID(db *gorm.DB, id uint16) _struct.Error {
	var entry _struct.Error

	db.Where("ErrorID = ?", id).Find(&entry)
	return entry
}

// Gets all entries from the database
func GetAllEntriesFromDB(db *gorm.DB) []_struct.Error {
	var errors []_struct.Error
	db.Find(&errors)
	fmt.Println(json.Marshal(errors))
	return errors
}
