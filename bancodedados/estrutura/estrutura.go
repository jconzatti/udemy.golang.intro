package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/microsoft/go-mssqldb"
)

func executar(db *sql.DB, sql string) sql.Result {
	resultado, erro := db.Exec(sql)
	if erro != nil {
		panic(erro)
	}
	return resultado
}

func main() {
	//servidor := "localhost\\express"
	//usuario := "ccj"
	//senha := "123mudar"
	//banco := "master"

	//parametrosParaConexao := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", servidor, usuario, senha, banco)

	//bancoDeDado, erro := sql.Open("sqlserver", parametrosParaConexao)
	bancoDeDado, erro := sql.Open("sqlite3", "./bancodedados/estrutura/teste.db3")
	if erro != nil {
		panic(erro)
	}
	defer bancoDeDado.Close()

	executar(bancoDeDado, "create table teste (id integer primary key, descricao text)")
	executar(bancoDeDado, "drop table teste")
}
