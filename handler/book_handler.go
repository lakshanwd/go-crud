package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

//BookGetHandler - handle book get requests
func BookGetHandler(w http.ResponseWriter, r *http.Request) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	books, err := bookRepo.Select()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convertListToArray(books))
}

//BookPostHandler - handle book post requests
func BookPostHandler(w http.ResponseWriter, r *http.Request) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}

	var book dao.Book
	json.NewDecoder(r.Body).Decode(&book)
	err = bookRepo.Insert(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

//BookPutHandler - handle book put requests
func BookPutHandler(w http.ResponseWriter, r *http.Request) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	var book dao.Book
	json.NewDecoder(r.Body).Decode(&book)
	err = bookRepo.Update(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

//BookDeleteHandler - handle book delete requests
func BookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	bookRepo, err := repository.GetBookRepository()
	defer bookRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}

	var book dao.Book
	book.BookID, _ = strconv.Atoi(mux.Vars(r)["id"])
	err = bookRepo.Remove(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
