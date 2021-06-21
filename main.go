package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Javascript", Preco: 39, Quantidade: 5},
		{"Tenis", "Olimpic", 89, 3},
		{"Fone", "Multilaser", 59, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
