package routes

import (
	"net/http"

	"github.com/gilsondev/produtosweb-go/controllers"
)

func CarregaRotas() {
  http.HandleFunc("/", controllers.Index)
  http.HandleFunc("/new", controllers.New)
  http.HandleFunc("/insert", controllers.Insert)
  http.HandleFunc("/edit", controllers.Edit)
  http.HandleFunc("/update", controllers.Update)
  http.HandleFunc("/remove", controllers.Remove)
}
