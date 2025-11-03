package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please send a GET request", http.StatusBadRequest)
		return
	}

	//parse query paramas
	query := r.URL.Query()
	// extract page no
	pageAsStr := query.Get("page")
	//extract data limit
	limitAsStr := query.Get("limit")

	//convert string to integer
	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	products, err := h.svc.List(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var cnt int64
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		count, _ := h.svc.Count()
		cnt = count
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		count, _ := h.svc.Count()
		cnt = count
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		count, _ := h.svc.Count()
		cnt = count
	}()

	utils.SendPage(w, products, page, limit, cnt)
}
