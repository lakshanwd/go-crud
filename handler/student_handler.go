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

//StudentPostHandler - handle student post request
func StudentPostHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	var student dao.Student
	json.NewDecoder(r.Body).Decode(student)
	err = studentRepo.Insert(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

//StudentGetHandler - handle student get request
func StudentGetHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	students, err := studentRepo.Select()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convertListToArray(students))
}

//StudentPutHandler - handle student put request
func StudentPutHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}

	var student dao.Student
	json.NewDecoder(r.Body).Decode(student)
	err = studentRepo.Update(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

//StudentDeleteHandler - handle student delete request
func StudentDeleteHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	var student dao.Student
	student.StudentID, _ = strconv.Atoi(mux.Vars(r)["id"])
	err = studentRepo.Remove(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}
