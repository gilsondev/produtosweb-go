package models

import "github.com/gilsondev/produtosweb-go/db"


type Product struct {
  Id int
  Nome string
  Descricao string
  Preco float64
  Quantidade int
}

func FetchAll() []Product {
  db := db.ConnectDb()
  defer db.Close()

  sql := "SELECT * FROM products"

  result, err := db.Query(sql)

  if err != nil {
    panic(err.Error())
  }

  p := Product{}
  produtos := []Product{}

  for result.Next() {
    var id, quantidade int
    var nome, descricao string
    var preco float64

    err := result.Scan(&id, &nome, &descricao, &preco, &quantidade)

    if err != nil {
      panic(err.Error())
    }

    p.Nome = nome
    p.Descricao = descricao
    p.Preco = preco
    p.Quantidade = quantidade

    produtos = append(produtos, p)
  }

  return produtos
}

func CreateNewProduct(nome string, descricao string, preco float64, quantidade int) {
  db := db.ConnectDb()
  defer db.Close()

  query, err := db.Prepare("INSERT INTO products (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")

  if err != nil {
    panic(err.Error())
  }

  query.Exec(nome, descricao, preco, quantidade)
}
