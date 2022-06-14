package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto {
		{Nome: "Mangá 'Remina' by Junji Ito", Descricao: "Número de páginas: 256; Idioma: Inglês; Editora: Viz Media", Preco: 112.93, Quantidade: 10},
		{"Mangá 'Smashed' by Junji Ito", "Número de páginas: 416; Idioma: Inglês; Editora: Viz Media", 201.50, 10},
		{"Mangá 'Frankenstein' by Junji Ito", "Número de páginas: 412; Idioma: Português; Editora: Pipoca e Nanquin", 52.40, 10},
		{"Mangá 'Tomie' by Junji Ito", "Número de páginas: 376; Idioma: Português; Editora: Pipoca e Nanquim", 97.86, 10},
		{"Mangá 'Uzumaki' by Junji Ito", "Número de páginas: 668; Idioma: Português; Editora: Pipoca e Nanquim", 146.59, 10},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
