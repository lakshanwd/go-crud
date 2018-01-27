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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}

	studentBooks, err := studentBookRepo.Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convertListToArray(studentBooks))
}

//StudentBookPostHandler - handle studentbook post requests
func StudentBookPostHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}

	var studentBook dao.StudentBook
	json.NewDecoder(r.Body).Decode(studentBook)
	err = studentBookRepo.Insert(studentBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

//StudentBookPutHandler - handle studentbook put requests
func StudentBookPutHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}

	var studentBook dao.StudentBook
	json.NewDecoder(r.Body).Decode(studentBook)
	err = studentBookRepo.Update(studentBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

//StudentBookDeleteHandler - handle studentbook delete requests
func StudentBookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	var studentBook dao.StudentBook
	err = studentBookRepo.Remove(studentBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
