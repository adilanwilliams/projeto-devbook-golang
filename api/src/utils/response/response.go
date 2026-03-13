package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response defines the standard structure for API responses.
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// ResponseJSON writes a JSON response to the client
// with the provided HTTP status code.
func ResponseJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if response.Success || response.Data != nil {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Fatal(err)
		}
	}
}

// ResponseError writes an error response in JSON format
// using the standard Response structure.
func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	ResponseJSON(w, statusCode, Response{
		Success: false,
		Data:    err.Error(),
	})
}
