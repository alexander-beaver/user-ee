package api

import (
	"fmt"
	"net/http"
)

func putAPIErrorHandler(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w,"Test")
}