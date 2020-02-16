package api

import (
	"fmt"
	"net/http"
)

// Handles a PUT call to the API for an Error
func PutErrorAPIHandler(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w,"Test")
}