package repositorios

import (
	models "api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios  {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error)  {
	query := "insert into usuarios (nome, nick, email, senha) values($1,$2,$3,$4) returning id"
	var ultimoIdInserido uint64

	erro := repositorio.db.QueryRow(query,usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&ultimoIdInserido)

	if erro != nil {
		return 0, erro
	}
	return ultimoIdInserido, nil
}

func (repositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error)  {
	nomeOuNick = fmt.Sprintf("%%s%%, nomeOuNick")

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE $1 or nick LIKE $2",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next(){
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil{
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}
	
	return usuarios, nil
}