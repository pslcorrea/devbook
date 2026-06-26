package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	models "api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CriarUsuario cria um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request)  {
	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return		
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)

		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
	}

	respostas.JSON(w, http.StatusCreated, usuario)

	// w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioId)))
}

//BuscarUsuarios encontra todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request)  {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}
//BuscarUsuario encontra um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10,64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusOK, usuario)
}
//AtualizarUsuario atualizar um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuário que não seja o seu."))
		return 
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return 
	}	

	var usuario models.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
//DeletarUsuario exclui um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é permitido excluir um usuário que não seja seu."))
		return 
	}	

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request)  {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
		return 
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()	

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request)  {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return 
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return 
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível parar de seguir você mesmo"))
		return 
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	defer db.Close()	

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}

	respostas.JSON(w, http.StatusNoContent, nil)	
}