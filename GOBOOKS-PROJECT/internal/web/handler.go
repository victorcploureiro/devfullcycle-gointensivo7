package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
)

type BookHandlers struct {
	service *service.BookService
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "failed to get Books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}