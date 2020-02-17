package api

import (
	"encoding/json"
	"github.com/alexander-beaver/user-ee/core"
	"github.com/alexander-beaver/user-ee/core/db"
	"github.com/alexander-beaver/user-ee/core/struct"

	"github.com/jinzhu/gorm"

	"net/http"
)

// Handles a PUT call to the API for an Error
func PutErrorAPIHandler(db2 *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var e _struct.Error
		err := decoder.Decode(e)
		if err != nil {
			panic(err)
		}
		db.WriteErrorToDB(db2, e)
		core.RespondOK(w, r, nil)

	}
}
