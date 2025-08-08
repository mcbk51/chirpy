package main

import (
	"net/http"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirpByID (w http.ResponseWriter, r *http.Request) {
  chirpIDStr := r.PathValue("chirpID")

	chirpID, err := uuid.Parse(chirpIDStr) 
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.dbQueries.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not retrieve chirp", err)
		return
	}
  
	respChirp := Chirp{
			ID:         chirp.ID,
			CreatedAt:  chirp.CreatedAt,
			UpdatedAt:  chirp.UpdatedAt,
			Body:       chirp.Body,
			UserID:     chirp.UserID,
		}
	respondWithJSON(w, http.StatusOK, respChirp)
}

func (cfg *apiConfig) handlerGetChirps (w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve chirps", err)
		return
	}

	respChirps := []Chirp{}  
	for _, chirp := range chirps {
		respChirps = append(respChirps, Chirp{
			ID:         chirp.ID,
			CreatedAt:  chirp.CreatedAt,
			UpdatedAt:  chirp.UpdatedAt,
			Body:       chirp.Body,
			UserID:     chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, respChirps)
}

