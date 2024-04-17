package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"raceia/db"
)

var robotDataList []db.RobotData

func PostDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var data db.RobotData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(data)
	robotDataList = append(robotDataList, data)

	db.InsertData(data)

	w.Write([]byte("Données reçues avec succès"))
}
