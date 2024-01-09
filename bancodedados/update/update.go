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

	comando, _ := bancoDeDado.Prepare("update usuarios set nome = ? where id = ?")
	comando.Exec("Jhoni Conzatti", 1)
	comando.Exec("Thaís de Lima Silva", 2)

	comando, _ = bancoDeDado.Prepare("delete from usuarios where id = ?")
	comando.Exec(4)
	comando.Exec(2000)
	comando.Exec(2001)
}
