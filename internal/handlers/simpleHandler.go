package handlers

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	response := struct {
		Message string `json:"message"`
	}{
		Message: "fuck u",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		w.Write(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	
	w.WriteHeader(http.StatusOK)
}
