package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/repository"
)

func main() {
	fmt.Println("handeling routes")
	http.HandleFunc("/student", handleStudent)
	http.HandleFunc("/book", handleBook)
	http.HandleFunc("/studentbook", handleStudentBook)

	fmt.Println("server listening on", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleStudent(w http.ResponseWriter, r *http.Request) {
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

func handleBook(w http.ResponseWriter, r *http.Request) {
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

func handleStudentBook(w http.ResponseWriter, r *http.Request) {
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
