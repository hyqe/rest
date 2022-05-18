package rest

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func ReadJson(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
