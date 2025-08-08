package main

import (
	"encoding/json"
	"net/http"
  "time"	

	"github.com/google/uuid"
	"github.com/mcbk51/chirpy/internal/auth"
)

type User struct {
		ID          uuid.UUID   `json:"id"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		Email       string      `json:"email"`
	  Password    string      `json:"-"`
}


func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request)  {
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
	  respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)	
		return 
	}

  hashedPass, err := auth.HashedPassword(reqBody.Password) 
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not hash password", err)
		return
	}
  
	user, err := cfg.db.CreateUser(r.Context(), reqBody.Email, hashedPass)
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
