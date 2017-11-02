package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/supunz/go-crud/handler"
)

func main() {
	fmt.Println("handeling routes")
	http.HandleFunc("/student", handler.StudentHandler)
	http.HandleFunc("/book", handler.BookHandler)
	http.HandleFunc("/studentbook", handler.StudentBookHandler)

	fmt.Println("server listening on", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
