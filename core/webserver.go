package core

import (
	"encoding/json"
	"net/http"
)

// A generic response that is sent back whenever the server needs to return data
type Response struct {
	Status uint16
	Body   interface{}
}

// Return a 200 (OK) response
func RespondOK(w http.ResponseWriter, r *http.Request, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := Response{
		Status: http.StatusOK,
		Body:   content,
	}

	json.NewEncoder(w).Encode(res)

}

// Return a 404 (Not Found) response
func RespondNotFound(w http.ResponseWriter, r *http.Request, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	res := Response{
		Status: http.StatusNotFound,
		Body:   content,
	}

	json.NewEncoder(w).Encode(res)

}

// Return a 500 (Internal Server Error) response
func RespondInternalServerError(w http.ResponseWriter, r *http.Request, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := Response{
		Status: http.StatusInternalServerError,
		Body:   content,
	}

	json.NewEncoder(w).Encode(res)

}
