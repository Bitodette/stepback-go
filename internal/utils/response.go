package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginatedResponse struct {
	Success  bool        `json:"success"`
	Count    int64       `json:"count"`
	Next     string      `json:"next,omitempty"`
	Previous string      `json:"previous,omitempty"`
	Data     interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, status int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func Message(w http.ResponseWriter, status int, message string) {
	JSON(w, status, Response{
		Success: true,
		Message: message,
	})
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, Response{
		Success: false,
		Message: message,
	})
}

func ValidationError(w http.ResponseWriter, errs interface{}) {
	JSON(w, http.StatusBadRequest, Response{
		Success: false,
		Message: "Validation failed",
		Errors:  errs,
	})
}

func Paginated(w http.ResponseWriter, count int64, next, previous string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PaginatedResponse{
		Success:  true,
		Count:    count,
		Next:     next,
		Previous: previous,
		Data:     data,
	})
}
