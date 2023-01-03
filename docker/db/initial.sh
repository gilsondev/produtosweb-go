#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE TABLE products (
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
  );

  INSERT INTO products (nome, descricao, preco, quantidade)
  VALUES 
  ('Camiseta', 'Preta', 19, 10),
  ('Fone', 'Legal', 99, 5);
EOSQL
