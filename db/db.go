package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=admin host=localhost sslmode=disable"

	db, erro := sql.Open("postgres", conexao)

	if erro != nil {
		panic(erro.Error())
	}

	return db
}
