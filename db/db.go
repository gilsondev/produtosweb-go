package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
  connStr := "postgres://postgres:postgres@localhost:5432/products_db?sslmode=disable"
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    panic(err.Error())
  }

  return db
}
