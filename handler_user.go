package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/marbh56/hypecharacter/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decode := json.NewDecoder(r.Body)
	params := parameters{}
	err := decode.Decode(&params)
	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))

	}

	respondWithJSON(w, 200, user)
}