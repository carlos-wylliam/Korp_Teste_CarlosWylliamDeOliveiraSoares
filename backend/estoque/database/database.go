package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Conectar() {
	db, err := sql.Open("sqlite", "./estoque.db")
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	criarTabela := `
	CREATE TABLE IF NOT EXISTS produtos (
		codigo TEXT PRIMARY KEY,
		descricao TEXT NOT NULL,
		saldo REAL NOT NULL
	);`

	_, err = db.Exec(criarTabela)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	DB = db
}