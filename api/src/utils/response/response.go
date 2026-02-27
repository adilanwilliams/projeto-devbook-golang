package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response represents a response for API.
type Response struct {
	Success bool   `json:"success"`
	Data interface{} `json:"data"`
}

// ResponseJSON returns a response JSON for the requests.
func ResponseJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil{
		log.Fatal(err)
	}
}

// ResponseJSON returns a response JSON for the errors on requests.
func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	ResponseJSON(w, statusCode, Response{
		Success: false,
		Data: err.Error(),
	})
}
