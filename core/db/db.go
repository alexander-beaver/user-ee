package db
import (
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




}

func GetEntryFromDBGivenID(db *gorm.DB, id uint16) _struct.Error {
	var entry _struct.Error



	db.Where("name = ?", id).Find(&entry)
	return entry
}