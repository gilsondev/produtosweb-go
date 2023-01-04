package main

import (
	"net/http"

	"github.com/gilsondev/produtosweb-go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
