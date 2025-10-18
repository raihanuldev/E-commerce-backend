package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var logginUser ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&logginUser)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Invaild Request", 400)
		return
	}
	// createdUSer := newUser.Store()
	logggedUser := database.Find(logginUser.Email)

	if logggedUser == nil {
		http.Error(w, "Invaild Createditional", http.StatusBadRequest)
		return
	}
	utils.SendData(w, logggedUser, http.StatusCreated)
}
