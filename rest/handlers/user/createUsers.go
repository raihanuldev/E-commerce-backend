package user

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	//catching data from request
	var newUser domain.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Invaild Request", 400)
		return
	}
	createdUSer, err := h.svc.Create(newUser)
	if err != nil {
		fmt.Print("error", err)
	}
	utils.SendData(w, createdUSer, http.StatusCreated)

}
