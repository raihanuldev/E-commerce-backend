package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Implement JWT Checking
			header := r.Header.Get("Authorization") //got it full JWT
			// fmt.Println("Header-----", header)
			headerArr := strings.Split(header, " ") //spilt from jwt for get Token
			if len(headerArr) != 2 {
				http.Error(w, "UnAuthorized", http.StatusUnauthorized)
				return
			}
			accessTokenFromClient := headerArr[1]
			// fmt.Println("Access token---------", accessToken)
			accessTokenArr := strings.Split(accessTokenFromClient, ".")
			if len(accessTokenArr) != 3 {
				http.Error(w, "UnAuthorized", http.StatusUnauthorized)
				return
			}
			//Hash create
			// cnf := config.GetConfig()  //convert it Layer

			/**
			header=> accessToken[0]
			payload=> accessToken[1]
			signure=> accessToken[2]
			*/
			message := accessTokenArr[0] + "." + accessTokenArr[1]

			// secretByteArr := []byte(cnf.SecretKey)
			secretByteArr := []byte(m.cnf.SecretKey)
			messagebyteArr := []byte(message)

			h := hmac.New(sha256.New, secretByteArr)
			h.Write(messagebyteArr)

			hash := h.Sum(nil)

			newSignure := Base64URLEncode(hash)

			if newSignure != accessTokenArr[2] {
				http.Error(w, "UnAuthrized, Tui Hacker Hala", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}

func Base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
