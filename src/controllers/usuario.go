package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//CriarUsuario cria um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request)  {
	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()

	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioId, erro := repositorio.Criar(usuario)

		if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioId)))
}

//BuscarUsuarios encontra todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Buscando todos os Usuários!"))
}
//BuscarUsuario encontra um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Buscndo um  Usuário!"))
}
//AtualizarUsuario atualizar um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Atualizando Usuário!"))
}
//DeletarUsuario exclui um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Excluindo Usuário!"))
}