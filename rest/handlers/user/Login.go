package user

import (
	"ecommerce/config"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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
	logggedUser := h.userRepo.Find(logginUser.Email)

	if logggedUser == nil {
		http.Error(w, "Invaild Createditional", http.StatusBadRequest)
		return
	}
	cnf := config.GetConfig()

	accessToken, err := utils.CreateJwt(cnf.SecretKey, utils.Payload{
		Sub:         logggedUser.ID,
		FristName:   logggedUser.FristName,
		LastName:    logggedUser.LastName,
		Email:       logggedUser.Email,
		IsShopOwner: logggedUser.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Internel Server Error", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, accessToken, http.StatusCreated)
}
