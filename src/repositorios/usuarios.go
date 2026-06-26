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
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome ILIKE $1 or nick ILIKE $2",
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

func (repositorio usuarios) BuscarPorID(ID uint64) (models.Usuario, error)  {
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = $1",
		ID,
	)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,			
		); erro != nil {
			return  models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Atualizar(ID uint64, usuario models.Usuario) error{
	_, erro := repositorio.db.Exec(
		"UPDATE usuarios SET nome = $1, nick = $2, email = $3 WHERE id = $4",
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		ID,
	)

	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) Deletar(ID uint64) error{
	_, erro := repositorio.db.Exec(
		"DELETE FROM usuarios WHERE id = $1",
		ID,
	)

	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error){
	var usuario models.Usuario
	erro := repositorio.db.QueryRow("SELECT id, senha FROM usuarios WHERE email = $1", email).Scan(&usuario.ID, &usuario.Senha)
	
	if erro != nil {
		return models.Usuario{}, erro
	}

	return usuario, nil
}

func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64)  error {
	query := "insert into seguidores (usuario_id, seguidor_id) values($1, $2) ON CONFLICT (usuario_id, seguidor_id) DO NOTHING"

	_, erro := repositorio.db.Exec(query, usuarioID, seguidorID )
	if erro != nil {
		return erro
	}

	return  nil
}

func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64)  error {
	query := "delete from seguidores where usuario_id = $1 and seguidor_id = $2"

	if _, erro := repositorio.db.Exec(query, usuarioID, seguidorID ); erro != nil {
		return erro
	}

	return  nil
}

func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error){
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = $1
	`, usuarioID)
	
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

func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error){
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = $1
	`, usuarioID)
	
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

