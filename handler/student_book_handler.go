package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

//StudentBookHandler - handle studentbook request
func StudentBookHandler(w http.ResponseWriter, r *http.Request) {
	studentBookRepo, err := repository.GetStudentBookRepository()
	defer studentBookRepo.Close()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		{
			studentBooks, err := studentBookRepo.Select()
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(studentBooks)
		}
	case http.MethodPost:
		{
			var studentBook dao.StudentBook
			json.NewDecoder(r.Body).Decode(studentBook)
			err := studentBookRepo.Insert(studentBook)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	case http.MethodPut:
		{
			var studentBook dao.StudentBook
			json.NewDecoder(r.Body).Decode(studentBook)
			err := studentBookRepo.Update(studentBook)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	case http.MethodDelete:
		{
			var studentBook dao.StudentBook
			json.NewDecoder(r.Body).Decode(studentBook)
			err := studentBookRepo.Remove(studentBook)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(true)
		}
	}
}
