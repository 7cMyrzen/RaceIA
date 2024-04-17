package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"raceia/db"
)

var raceDataList []db.RaceData

func PostRaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var data db.RaceData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(data)
	raceDataList = append(raceDataList, data)

	w.Write([]byte("Données reçues avec succès"))
}
