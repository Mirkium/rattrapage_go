package templates

import (
	"fmt"
	"html/template"
)

var Temp *template.Template

func InitTemplates() {
	temp, err := template.ParseGlob("./*.html")
	if (err != nil) {
		fmt.Println("Erreur dans la récupération des templates : ", err)
	}
	Temp = temp
}