package api

import (
	"encoding/json"
	"github.com/alexander-beaver/user-ee/core"
	"github.com/alexander-beaver/user-ee/core/db"
	"github.com/alexander-beaver/user-ee/core/struct"

	"github.com/jinzhu/gorm"

	"net/http"
)

// The body that is expected in a JSON request
// An example request is:
//curl --request POST \
//  --url http://localhost:8080/write \
//  --header 'content-type: application/json' \
//  --data '{
//      "EndpointID": "123",
//	"ErrorID": 100,
//	"ErrorMessage": "Test",
//	"Comments": "Test"
//}'
type BodyJSON struct {
	EndpointID   string
	ErrorID      uint16
	ErrorMessage string
	Comments     string
}

// Handles a PUT call to the API for an Error
// Endpoint: http://localhost:8080/write
func PutErrorAPIHandler(db2 *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var e BodyJSON
		err := decoder.Decode(&e)
		if err != nil {
			core.RespondInternalServerError(w, r, err)
			panic(err)
		}
		writeStruct := _struct.Error{
			EndpointID:   e.EndpointID,
			ErrorID:      e.ErrorID,
			ErrorMessage: e.ErrorMessage,
			Comments:     e.Comments,
		}
		db.WriteErrorToDB(db2, writeStruct)
		core.RespondOK(w, r, nil)

	}
}
