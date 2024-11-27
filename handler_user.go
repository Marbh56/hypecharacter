package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/marbh56/hypecharacter/internal/Auth"
	"github.com/marbh56/hypecharacter/internal/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error hashing password: %v", err))
		return
	}

	// Create the user in the database
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: string(hashedPassword),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	email, password, err := Auth.GetEmailAndPassword(r.Header)

	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Error getting credentials: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), email)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		respondWithError(w, 401, "Invalid password")
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}
