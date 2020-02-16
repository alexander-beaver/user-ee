package db
import (
	"github.com/alexander-beaver/user-ee/core/struct"

	"github.com/jinzhu/gorm"
	"time"
)

var dbName = "test.db"


func SetupDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	db.Set("gorm:auto_preload", true)

	return db
}


func writeErrorToDB(reported Error){



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