package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/db"
	"gopkg.in/mgo.v2/bson"
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
	db, err := db.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Session.Close()
	studentCollection := db.C("student")

	switch r.Method {

	//handle get requests
	case http.MethodGet:
		{
			var result []dao.Student
			err := studentCollection.Find(nil).All(&result)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			json.NewEncoder(w).Encode(result)
		}

	//handle post requests
	case http.MethodPost:
		{
			var student dao.Student
			json.NewDecoder(r.Body).Decode(student)
			err := studentCollection.Insert(student)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

		}

	//handle delete requests
	case http.MethodDelete:
		{
			err := studentCollection.Remove(bson.M{"Id": r.FormValue("studentId")})
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
		}
	}
}

func handleBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(dao.Book{})
	}
}

func handleStudentBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		st := dao.Student{}
		books := []dao.Book{}
		json.NewEncoder(w).Encode(dao.StudentBook{Student: &st, Books: &books})
	}
}
