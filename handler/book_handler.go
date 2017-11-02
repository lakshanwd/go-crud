package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

//BookHandler - handle book request
func BookHandler(w http.ResponseWriter, r *http.Request) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		{
			books, err := bookRepo.Select()
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(books)
		}
	case http.MethodPut:
		{
			var book dao.Book
			json.NewDecoder(r.Body).Decode(book)
			err := bookRepo.Update(book)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	case http.MethodPost:
		{
			var book dao.Book
			json.NewDecoder(r.Body).Decode(book)
			err := bookRepo.Insert(book)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	case http.MethodDelete:
		{
			var book dao.Book
			json.NewDecoder(r.Body).Decode(book)
			err := bookRepo.Update(book)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	}
}
