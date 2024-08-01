package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dnreikronos/wiki-fans-backend/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var character models.Character

	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(character)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Character inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
