package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productid := r.PathValue("productId")
	fmt.Println(productid)
	pid, err := strconv.Atoi(productid)
	if err != nil {
		http.Error(w, "please give Valid ID", 400)
		return
	}
	product := h.productRepo.Get(pid)
	if product == nil {
		utils.SendData(w, "Prodcut Not found", 404) //TODO: Create a Function for SendError
	} else {
		utils.SendData(w, product, 200)
	}
}
