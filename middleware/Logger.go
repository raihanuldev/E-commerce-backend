package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	ctrl := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println("Logger Info: ", "Method: ", r.Method, "URL: ", r.URL.Path, "Excution Times: ", time.Since(start))

	}
	return http.HandlerFunc(ctrl)
}
