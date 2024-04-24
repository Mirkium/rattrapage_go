package main

import (
	"fmt"
	"main/controller"
	"main/templates"
	"net/http"
)

func main() {
	templates.InitTemplates()
	
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/accueil", controller.Accueil)
	http.HandleFunc("/add", controller.Add)
	http.HandleFunc("/add_handle", controller.AddHandle)
	http.HandleFunc("/produit", controller.Produit)

	fmt.Println("http://localhost:8080/accueil")
	http.ListenAndServe("http://localhost:8080", nil)
}
