package middileware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	ctrl := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Logger Info: ", "Method: ", r.Method, "URL: ", r.URL.Path, "Excution Times: ", time.Since(start))
		log.Println("w: ", w)
		log.Println("r: ", r)
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(ctrl)
}
