package response

import (
	"encoding/json"
	"net/http"
	"time"
)

// ApiResponse represents the generic structure for all API responses.
type ApiResponse[T any] struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Data      T      `json:"data"`
	Path      string `json:"path"`
}

// New creates a new structured response payload.
func New[T any](status int, success bool, message string, data T, path string) ApiResponse[T] {
	return ApiResponse[T]{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Status:    status,
		Success:   success,
		Message:   message,
		Data:      data,
		Path:      path,
	}
}

// Send writes the JSON response to the client.
func Send[T any](w http.ResponseWriter, status int, success bool, message string, data T, path string) {
	response := New(status, success, message, data, path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}
