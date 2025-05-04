package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gaashvik/Go--Learning-repo/project/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:name`
	}
	decoder := json.NewDecoder(r.Body) //returns a decoder
	param := params{}
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json:%s", err))
		return
	}
	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: param.Name})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create user:%s", err))
		return
	}

	respondWithJson(w, 200, user)
}
