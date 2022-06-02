package handlers

import (
	"API-GO-com-postgre/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removido %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Registro removido com sucesso",
	}
	w.Header().Add("Cpntent-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
