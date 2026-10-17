package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status int, data any) {

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, data any){
	 Response(w , status , map[string]any{
		"Error": data,
	 })
}
