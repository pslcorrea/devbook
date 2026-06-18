package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexao é a string de conexão com Postgres
	StringConexao = ""
	Porta         = 0
	// Chave para assinar o token
	Scretkey []byte
)

// Carregar vai iniciar as váriaveis de ambiente
func Carregar() {
	var erro error

	if erro := godotenv.Load(); erro!= nil{
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}
	dbHost := os.Getenv("DB_HOST")
	dbUsuario := os.Getenv("DB_USUARIO")
	dbSenha := os.Getenv("DB_SENHA")
	dbNome := os.Getenv("DB_NOME")
	dbPorta := os.Getenv("DB_PORT")

	if dbPorta == "" {
		dbPorta = "5432"
	}

	StringConexao = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsuario, dbSenha, dbNome, dbPorta,
	)

	Scretkey = []byte(os.Getenv("SECRET_KEY"))

}