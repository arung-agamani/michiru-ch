package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool `json:"status"`
	Error []string `json:"error"`
	Data interface{} `json:"data"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, success bool, errors []string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: success,
		Error: errors,
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteSuccessJSON(w http.ResponseWriter, data interface{}) {
	WriteJSONResponse(w, http.StatusOK, true, []string{}, data)
}

func WriteBadRequestJSON(w http.ResponseWriter, errors []string) {
	WriteJSONResponse(w, http.StatusBadRequest, false, errors, nil)
}

func WriteInternalServerErrorJSON(w http.ResponseWriter, errors []string) {
	WriteJSONResponse(w, http.StatusInternalServerError, false, errors, nil)
}
