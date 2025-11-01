package utils

import "net/http"

type PaginatedData struct {
	Data        any         `json:"data"`
	Paginations Pagniations `json:"paginations"`
}

type Pagniations struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, count int64) {
	paginated := PaginatedData{
		Data: data,
		Paginations: Pagniations{

			Page:       page,
			Limit:      limit,
			TotalItems: count,
			TotalPages: count / limit,
		},
	}
	SendData(w, paginated, 200)
}
