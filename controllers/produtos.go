package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gilsondev/produtosweb-go/models"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
  produtos := models.FetchAll()

  templates.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
  templates.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    nome := r.FormValue("nome")
    descricao := r.FormValue("descricao")
    precoRaw := r.FormValue("preco")
    quantidadeRaw := r.FormValue("quantidade")

    preco, err := strconv.ParseFloat(precoRaw, 64)

    if err != nil {
      log.Println("Erro na conversão do preço: ", err)
    }

    quantidade, err := strconv.Atoi(quantidadeRaw)

    if err != nil {
      log.Println("Erro na conversão da quantidade: ", err)
    }

    models.CreateNewProduct(nome, descricao, preco, quantidade)

    http.Redirect(w, r, "/", 301)
  }
}

func Edit(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")

  produto := models.Fetch(id)

  templates.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    id := r.FormValue("id")
    nome := r.FormValue("nome")
    descricao := r.FormValue("descricao")
    precoRaw := r.FormValue("preco")
    quantidadeRaw := r.FormValue("quantidade")

    preco, err := strconv.ParseFloat(precoRaw, 64)

    if err != nil {
      log.Println("Erro na conversão do preço: ", err)
    }

    quantidade, err := strconv.Atoi(quantidadeRaw)

    if err != nil {
      log.Println("Erro na conversão da quantidade: ", err)
    }

    models.UpdateProduct(id, nome, descricao, preco, quantidade)

    http.Redirect(w, r, "/", 301)
  }
}

func Remove(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")

  models.RemoveProduct(id)

  http.Redirect(w, r, "/", 301)
}
