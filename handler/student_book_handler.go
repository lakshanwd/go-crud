package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

//StudentBookGetHandler - handle studentbook get requests
func StudentBookGetHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	studentBooks, err := studentBookRepo.Select()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(studentBooks)
}

//StudentBookPostHandler - handle studentbook post requests
func StudentBookPostHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var studentBook dao.StudentBook
	json.NewDecoder(r.Body).Decode(studentBook)
	err = studentBookRepo.Insert(studentBook)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(true)
}

//StudentBookPutHandler - handle studentbook put requests
func StudentBookPutHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var studentBook dao.StudentBook
	json.NewDecoder(r.Body).Decode(studentBook)
	err = studentBookRepo.Update(studentBook)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(true)
}

//StudentBookDeleteHandler - handle studentbook delete requests
func StudentBookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var studentBook dao.StudentBook
	json.NewDecoder(r.Body).Decode(studentBook)
	err = studentBookRepo.Remove(studentBook)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(true)
}
