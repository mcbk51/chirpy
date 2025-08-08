package main

import (
	"encoding/json"
	"net/http"
  "time"	

	"github.com/google/uuid"
	"github.com/mcbk51/chirpy/internal/auth"
)


func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request)  {
	
	type requestBody struct {
		Password string `json:"password"`
	  Email    string `json:"email"`	
	}

	type response struct {
	  User	
	}

  decoder := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
	  respondWithError(w, http.StatusBadRequest, "Could not decode parameters", err)	
		return 
	}

  user, err := cfg.db.getUserByEmail(r.Context(), reqBody.Email)
	if err != nil {
	  respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)	
		return 
	}
	
	err = auth.CheckPasswordHash(reqBody.Password, user.Password) 		

	if err != nil {
	  respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)	
		
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})

}
  
