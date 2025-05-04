package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (apiCFG *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:name`
	}
	decoder := json.NewDecoder(r.Body)
	param := params{}
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json:%s", err))
	}
	user, err := apiCFG.DB.GetUser(r.Context(), param.Name)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not find user%s", err))
		return
	}
	respondWithJson(w, 200, user)

}
