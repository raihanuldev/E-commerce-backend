package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Create Product
func CreateProduct(w http.ResponseWriter, r *http.Request) {

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
	cnf := config.GetConfig()
	/**
	header=> accessToken[0]
	payload=> accessToken[1]
	signure=> accessToken[2]
	*/
	message := accessTokenArr[0] + "." + accessTokenArr[1]

	secretByteArr := []byte(cnf.SecretKey)
	messagebyteArr := []byte(message)

	h := hmac.New(sha256.New, secretByteArr)
	h.Write(messagebyteArr)

	hash := h.Sum(nil)

	newSignure := Base64URLEncode(hash)

	if newSignure != accessTokenArr[2] {
		http.Error(w, "UnAuthrized, Tui Hacker Hala", http.StatusUnauthorized)
		return
	}

	//cheking if request are vaild or invaild
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}

	database.Store(newProduct)
	utils.SendData(w, newProduct, 201)

}

func Base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
