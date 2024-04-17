package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// BoardHandler est le gestionnaire pour la route /board
func BoardHandler(w http.ResponseWriter, r *http.Request) {
	boardFilePath := "templates/board2.html"
	boardFile, err := template.ParseFiles(boardFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = boardFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
