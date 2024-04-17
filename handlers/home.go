package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// HomeHandler est le gestionnaire pour la route /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeFilePath := "templates/home.html"
	homeFile, err := template.ParseFiles(homeFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = homeFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
