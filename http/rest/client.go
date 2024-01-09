package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	idNaURL := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(idNaURL)
	switch {
	case r.Method == "GET" && id > 0:
		obterUsuarioPorID(w, r, id)
	case r.Method == "GET":
		obterUsuarios(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func obterUsuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	bancoDeDado, erro := sql.Open("sqlite3", "C:/Projetos Pessoais/udemy/golang/curso/bancodedados/teste.db3")
	if erro != nil {
		log.Fatal(erro)
	}
	defer bancoDeDado.Close()

	dadoDeUsuario := bancoDeDado.QueryRow("select id, nome from usuarios where id = ?", id)
	var usuario Usuario
	dadoDeUsuario.Scan(&usuario.ID, &usuario.Nome)
	jsonDeUsuario, _ := json.Marshal(usuario)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonDeUsuario))
}

func obterUsuarios(w http.ResponseWriter, r *http.Request) {
	bancoDeDado, erro := sql.Open("sqlite3", "C:/Projetos Pessoais/udemy/golang/curso/bancodedados/teste.db3")
	if erro != nil {
		log.Fatal(erro)
	}
	defer bancoDeDado.Close()

	dadosDeUsuarios, _ := bancoDeDado.Query("select id, nome from usuarios")
	defer dadosDeUsuarios.Close()

	var usuarios []Usuario
	for dadosDeUsuarios.Next() {
		var usuario Usuario
		dadosDeUsuarios.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	jsonDeUsuarios, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonDeUsuarios))
}
