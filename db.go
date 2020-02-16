package main
import (
	"github.com/jinzhu/gorm"
	"time"
)

var dbName = "test.db"
type Error struct{
	gorm.Model

	EndpointID string
	ErrorID uint16
	ErrorMessage string
	Comments string
	Time time.Time
}


func writeErrorToDB(reported Error){

	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(reported)

}

func getEntryFromDB(id uint16) Error {
	var entry Error
	db, err := gorm.Open("sqlite3", dbName)
	if(err != nil){
		panic("failed to connect database")

	}

	db.Where("name = ?", id).Find(&entry)
	return entry
}