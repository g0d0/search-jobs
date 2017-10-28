package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	vagas []Vaga
}

func (s *Storage) list() []Vaga {
	return s.vagas
}

func (s *Storage) add(v Vaga) {
	s.vagas = append(s.vagas, v)
}

func (s *Storage) save(v Vaga) {
	database, _ := sql.Open("sqlite3", "./vagas.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS vagas (id INTEGER PRIMARY KEY, data DATE, nome TEXT, cidade TEXT, url TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO vagas (data, nome, cidade, url) values(?,?,?,?)")
	statement.Exec(v.data, v.nome, v.cidade, v.url)
}