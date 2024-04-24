package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"main/templates"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Tailles string `json:"tailles"`
	ImgLink string `json:"img"`
	Reduc int `json:"reduc"`
}

func Accueil(w http.ResponseWriter, r *http.Request) {
	Data := GetAllProducts()

	templates.Temp.ExecuteTemplate(w, "accueil", Data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "add", nil)
}

func AddHandle(w http.ResponseWriter, r *http.Request) {
	var Data Product
	Data.Name = r.PostFormValue("name")
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Erreur dans la récupération de l'image : ", err)
	}
	defer file.Close()

	filepath := "./assets/img/" + handler.Filename
	f, _ := os.Create(filepath)
	defer f.Close()
	io.Copy(f, file)

	Data.ImgLink = filepath
	Data.Price, _ = strconv.Atoi(r.PostFormValue("price"))
	Data.Reduc, _ = strconv.Atoi(r.PostFormValue("reduc"))
	Data.Tailles = r.PostFormValue("tailles")

	Data.Id = len(GetAllProducts())

	Json, _ := os.ReadFile("./data.json")
	var JsonData []Product
	json.Unmarshal(Json, &JsonData)
	
	JsonData = append(JsonData, Data)
	
	data, err := json.MarshalIndent(JsonData, "", "  ")
	if err != nil {
		fmt.Println("Erreur dans la conversion des données en JSON : ", err)
	}

	os.WriteFile("./data.json", data, 0644)

	link := "/produit?id=" + strconv.Itoa(Data.Id)
	http.Redirect(w, r, link, http.StatusSeeOther)
}

func Produit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	Data := GetProduct(id)

	templates.Temp.ExecuteTemplate(w, "produit", Data)
}

///////////////////////////////////////////////////////////////////////////////////////

func GetAllProducts() []Product {
	products, err := os.ReadFile("./data.json")
	if err != nil {
		fmt.Println("Erreur dans la lecture du fichier data : ", err)
	}

	var AllProducts []Product

	json.Unmarshal(products, &AllProducts)

	return AllProducts
}

func GetProduct(id int) Product {
	products := GetAllProducts()
	for p := 0; p < len(products); p++ {
		if products[p].Id == id {
			return products[p]
		}
	}
	return Product{}
}