package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status uint16
	Body   interface{}
}

func RespondOK(w http.ResponseWriter, r *http.Request, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := Response{
		Status: http.StatusOK,
		Body:   content,
	}

	json.NewEncoder(w).Encode(res)

}
func RespondInternalServerError(w http.ResponseWriter, r *http.Request, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := Response{
		Status: http.StatusInternalServerError,
		Body:   content,
	}

	json.NewEncoder(w).Encode(res)

}
