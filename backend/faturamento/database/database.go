package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Conectar() {
	db, err := sql.Open("sqlite", "./faturamento.db")
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	criarTabelaNotas := `
	CREATE TABLE IF NOT EXISTS notas (
		numero INTEGER PRIMARY KEY AUTOINCREMENT,
		status TEXT NOT NULL
	);`

	criarTabelaItens := `
	CREATE TABLE IF NOT EXISTS itens_nota (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nota_numero INTEGER NOT NULL,
		produto_codigo TEXT NOT NULL,
		quantidade REAL NOT NULL,
		FOREIGN KEY (nota_numero) REFERENCES notas(numero)
	);`

	_, err = db.Exec(criarTabelaNotas)
	if err != nil {
		log.Fatal("Erro ao criar tabela notas:", err)
	}

	_, err = db.Exec(criarTabelaItens)
	if err != nil {
		log.Fatal("Erro ao criar tabela itens_nota:", err)
	}

	DB = db
}