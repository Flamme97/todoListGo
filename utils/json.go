package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string){

	if code > 499 {
		log.Println("responded with 5xx Error: ", msg)

	}
	type errResponse struct {
		Error string `json:"error"`
	} 
	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})

}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshel JSON response: %v", payload)
		w.WriteHeader(500)
		return 
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}