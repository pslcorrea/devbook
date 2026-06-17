package seguranca

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerificaSenha(SenhaComHash, senhaString  string) error  {
	return bcrypt.CompareHashAndPassword([]byte(SenhaComHash), []byte(senhaString))
}