package models

import "github.com/oklog/ulid/v2"

type Student struct {
	Name string `json:"nome_do_aluno"`
	ID   string `json:"id_do_aluno"`
}

var Alunos []Student = []Student{
	{
		Name: "Bertram Gilfoyle",
		ID:   ulid.Make().String(),
	},
	{
		Name: "Dinesh",
		ID:   ulid.Make().String(),
	},
}
