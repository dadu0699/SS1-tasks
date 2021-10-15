package app

import (
	"api-go/pkg/api"
	"encoding/json"
	"log"
	"net/http"
)

// parse decodes the content received in the body of a request to an interface.
func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// sendResponse send the status code and content encoded in json format to a request.
func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	resp := api.Response{
		Code: statusCode,
		Data: data,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}
