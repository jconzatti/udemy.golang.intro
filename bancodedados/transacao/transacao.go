package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	bancoDeDado, erro := sql.Open("sqlite3", "./bancodedados/teste.db3")
	if erro != nil {
		log.Fatal(erro)
	}
	defer bancoDeDado.Close()

	transacao, _ := bancoDeDado.Begin()
	comando, _ := transacao.Prepare("insert into usuarios (id, nome) values (?, ?)")
	comando.Exec(4000, "Bia")
	comando.Exec(4001, "Carlos")
	_, erro = comando.Exec(1, "Tiago")
	if erro != nil {
		transacao.Rollback()
		log.Fatal(erro)
	}
	transacao.Commit()
}
