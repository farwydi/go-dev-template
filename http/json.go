package dev_tempalte_http

import (
	"encoding/json"
	"net/http"
)

func ToJSON200(w http.ResponseWriter, s interface{}) {
	if b, err := json.Marshal(s); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}

	http.Error(w, "Bad JSON encode", http.StatusInternalServerError)
}

func ToJSON(w http.ResponseWriter, status int, s interface{}) {
	if b, err := json.Marshal(s); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(b)
		return
	}

	http.Error(w, "Bad JSON encode", http.StatusInternalServerError)
}
