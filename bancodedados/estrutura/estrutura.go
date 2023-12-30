package main

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func executar(db *sql.DB, sql string) sql.Result {
	resultado, erro := db.Exec(sql)
	if erro != nil {
		panic(erro)
	}
	return resultado
}

func main() {
	servidor := "localhost\\express"
	usuario := "ccj"
	senha := "123mudar"
	banco := "master"

	parametrosParaConexao := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", servidor, usuario, senha, banco)

	bancoDeDado, erro := sql.Open("sqlserver", parametrosParaConexao)
	if erro != nil {
		panic(erro)
	}
	defer bancoDeDado.Close()

	executar(bancoDeDado, "create database cod3rcursogo")
}
