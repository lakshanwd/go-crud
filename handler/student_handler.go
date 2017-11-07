package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

//StudentPostHandler - handle student post request
func StudentPostHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var student dao.Student
	json.NewDecoder(r.Body).Decode(student)
	err = studentRepo.Insert(student)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(true)
}

//StudentGetHandler - handle student get request
func StudentGetHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	students, err := studentRepo.Select()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

//StudentPutHandler - handle student put request
func StudentPutHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var student dao.Student
	json.NewDecoder(r.Body).Decode(student)
	err = studentRepo.Update(student)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(true)
}

//StudentDeleteHandler - handle student delete request
func StudentDeleteHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var student dao.Student
	json.NewDecoder(r.Body).Decode(student)
	err = studentRepo.Remove(student)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(true)
}
