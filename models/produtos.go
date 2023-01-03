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

    p.Id = id
    p.Nome = nome
    p.Descricao = descricao
    p.Preco = preco
    p.Quantidade = quantidade

    produtos = append(produtos, p)
  }

  return produtos
}

func Fetch(id string) Product {
  db := db.ConnectDb()
  query, err := db.Query("SELECT * FROM products WHERE id = $1 ORDER BY id ASC", id)

  if err != nil {
    panic(err.Error())
  }

  var product Product

  for query.Next() {
    err := query.Scan(&product.Id, &product.Nome, &product.Descricao, &product.Preco, &product.Quantidade)

    if err != nil {
      panic(err.Error())
    }
  }

  return product
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

func UpdateProduct(id string, nome string, descricao string, preco float64, quantidade int) {
  db := db.ConnectDb()
  defer db.Close()

  query, err := db.Prepare("UPDATE products SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")

  if err != nil {
    panic(err.Error())
  }

  query.Exec(nome, descricao, preco, quantidade, id)
}

func RemoveProduct(id string) {
  db := db.ConnectDb()
  defer db.Close()

  query, err := db.Prepare("DELETE FROM products WHERE id = $1")

  if err != nil {
    panic(err.Error())
  }

  query.Exec(id)
}
