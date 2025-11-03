package order

import (
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type reqBody struct {
	OrderID int64  `json:"orderId"`
	Status  string `json:"status"`
}

func (h *Handler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Please Sent PUT Request", 400)
		return
	}
	var reqdata reqBody
	// orderID int64, newStatus string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqdata)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}
	isUpdated, err := h.svc.UpdateOrderStatus(reqdata.OrderID, reqdata.Status)
	if err != nil {
		http.Error(w, "Problem for Updateing Statrus", 400)
	}
	if isUpdated == 1 {
		fmt.Println(isUpdated)
		utils.SendData(w, "Status Changed Succesfully", http.StatusCreated)
	}

}
