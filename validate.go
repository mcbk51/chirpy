package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Body string `json:"body"`
	}
	type validResponse struct {
		Valid bool `json:"valid"`
	}
	
	decoder := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}
	
	if len(reqBody.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	
	respondWithJSON(w, http.StatusOK, validResponse{
		Valid: true,
	})
}
