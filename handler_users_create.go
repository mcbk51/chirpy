package main

import (
	"encoding/json"
	"net/http"
  "time"	

	"github.com/google/uuid"
)

type User struct {
		ID          uuid.UUID   `json:"id"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		Email       string      `json:"email"`
}


func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request)  {
	type requestBody struct {
	  Email string `json:"email"`	
	}

	type response struct {
		User
	}

  decoder := json.NewDecoder(r.Body)
	reqBody := requestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
	  respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)	
		return 
	}

  
	user, err := cfg.dbQueries.CreateUser(r.Context(), reqBody.Email)
	if err != nil {
	  respondWithError(w, http.StatusInternalServerError, "Could not create user", err)	
		return 
	}
  
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
				ID:             user.ID,
				CreatedAt:      user.CreatedAt,
				UpdatedAt:      user.UpdatedAt,
				Email:          user.Email, 
		},
	})
}
