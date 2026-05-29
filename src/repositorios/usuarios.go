package repositorios

import (
	"api/src/modelos"
	"database/sql"
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
	// if erro != nil {
	// 	return 0, erro
	// }

	// defer repositorio.Close()

	erro := repositorio.db.QueryRow(query,usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&ultimoIdInserido)

	if erro != nil {
		return 0, erro
	}

	// ultimoIdInserido, erro := resultado.LastInsertId()

	// 	if erro != nil {
	// 	return 0, erro
	// }

	return ultimoIdInserido, nil
}