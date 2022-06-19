package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error writes the error to the response and logs it.
func Error(w http.ResponseWriter, r *http.Request, errorMessage string, statusCode int) {
	fmt.Println("err", errorMessage)

	jsonBytes, err := json.Marshal(struct {
		Error string
	}{errorMessage})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorMessage))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
