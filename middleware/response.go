package middleware

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

}
