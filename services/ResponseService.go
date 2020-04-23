package services

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	log.Println("<services ResponseWithJson func>")
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
