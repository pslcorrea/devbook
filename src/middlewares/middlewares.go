package middlewares

import (
	"fmt"
	"net/http"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando...")
		next(w, r)
	}
}