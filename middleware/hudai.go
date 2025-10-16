package middleware 

import (
	"fmt"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I'm from Hudai middleware!")
		next.ServeHTTP(w, r)
	})
}
