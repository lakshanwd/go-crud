package main

import "net/http"
import "encoding/json"

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type book struct {
	BookName string `json:"bookName"`
	Author   string `json:"author"`
}

type studentBook struct {
	Student *student `json:"student"`
	Books   *[]book  `json:"books"`
}

func main() {
	http.HandleFunc("/student", handleStudent)
	http.HandleFunc("/book", handleBook)
	http.HandleFunc("/studentbook", handleStudentBook)
	http.ListenAndServe(":8080", nil)
}

func handleStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(student{})
	}
}

func handleBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(book{})
	}
}

func handleStudentBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		student := student{}
		books := []book{book{}, book{}}
		json.NewEncoder(w).Encode(studentBook{Student: &student, Books: &books})
	}
}
