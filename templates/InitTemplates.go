package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplates() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erreur lors de la récupération du répertoire de travail : ", err)
		return
	}

	// Utilisez le chemin relatif pour rechercher les fichiers HTML dans le répertoire courant
	pattern := dir + "/templates/*.html"
	temp, err := template.ParseGlob(pattern)
	if err != nil {
		fmt.Println("Erreur dans la récupération des templates : ", err)
	}
	Temp = temp
}
