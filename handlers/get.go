package handlers

import (
	"encoding/json"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	dataJSON, err := json.Marshal(robotDataList)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}
