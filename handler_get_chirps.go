package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerGetChirps (w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve chirps", err)
		return
	}

	responseChirps := make([]Chirp, len(chirps))
	for i, chirp := range chirps {
		responseChirps[i] = Chirp {
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt:  chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}
	respondWithJSON(w, http.StatusOK, responseChirps)
}
