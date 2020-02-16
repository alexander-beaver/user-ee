package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct{
	Status uint16
	Body interface{}
}

func RespondOK(w http.ResponseWriter, r *http.Request, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := Response{
		Status: http.StatusOK,
		Body:   content,
	}

	json, err := json.Marshal(res)
	if(err != nil){
		panic(err)
	}
	fmt.Fprintf(w, string(json))

}
