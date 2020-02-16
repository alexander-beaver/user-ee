package db
import (
	"fmt"
	"github.com/alexander-beaver/user-ee/core/struct"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)

var dbName = "test.db"


func SetupDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&_struct.Error{})
	db.Set("gorm:auto_preload", true)

	return db
}


func WriteErrorToDB(db *gorm.DB, reported _struct.Error){
	db.Create(&reported)
	fmt.Println("Created Object")



}

func GetEntryFromDBGivenID(db *gorm.DB, id uint16) _struct.Error {
	var entry _struct.Error



	db.Where("ErrorID = ?", id).Find(&entry)
	return entry
}

func GetAllEntriesFromDB(db *gorm.DB) []_struct.Error{
	var errors []_struct.Error
	db.Find(&errors)
	fmt.Println("{}", errors)
	return errors
}