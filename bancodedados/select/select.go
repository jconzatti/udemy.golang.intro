package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	bancoDeDado, erro := sql.Open("sqlite3", "./bancodedados/teste.db3")
	if erro != nil {
		log.Fatal(erro)
	}
	defer bancoDeDado.Close()

	usuarios, _ := bancoDeDado.Query("select id, nome from usuarios")
	defer usuarios.Close()

	for usuarios.Next() {
		var u usuario
		usuarios.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
