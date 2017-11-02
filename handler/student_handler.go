package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

//StudentHandler - handle student request
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	studentRepo, err := repository.GetStudentRepository()
	defer studentRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Method {

	//handle get requests
	case http.MethodGet:
		{
			students, err := studentRepo.Select()
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(students)
		}

		//handle post requests
	case http.MethodPost:
		{
			var student dao.Student
			json.NewDecoder(r.Body).Decode(student)
			err := studentRepo.Insert(student)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}

		//handle delete requests
	case http.MethodDelete:
		{
			var student dao.Student
			json.NewDecoder(r.Body).Decode(student)
			err := studentRepo.Remove(student)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}

		//handle put requests
	case http.MethodPut:
		{
			var student dao.Student
			json.NewDecoder(r.Body).Decode(student)
			err := studentRepo.Update(student)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	}
}
