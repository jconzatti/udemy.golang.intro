package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	bancoDeDado, erro := sql.Open("sqlite3", "./bancodedados/teste.db3")
	if erro != nil {
		panic(erro)
	}
	defer bancoDeDado.Close()

	comando, _ := bancoDeDado.Prepare("insert into usuarios (nome) values (?)")
	comando.Exec("Maria")
	comando.Exec("João")
	resultado, _ := comando.Exec("Pedro")
	id, _ := resultado.LastInsertId()
	fmt.Println("O ID do Pedro é", id)

	linhasAfetadas, _ := resultado.RowsAffected()
	fmt.Println("O comando afetou", linhasAfetadas, "no banco de dados!")
}
