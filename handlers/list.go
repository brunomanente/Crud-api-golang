package handlers

import (
	"API-GO-com-postgre/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Panicf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Cpntent-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
