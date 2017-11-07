package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supunz/go-crud/handler"
)

func main() {
	fmt.Println("handeling routes")
	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).Path("/student").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPost).Path("/student").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPut).Path("/student").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodDelete).Path("/student").HandlerFunc(handler.StudentPostHandler)

	router.Methods(http.MethodGet).Path("/book").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPost).Path("/book").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPut).Path("/book").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodDelete).Path("/book").HandlerFunc(handler.StudentPostHandler)

	router.Methods(http.MethodGet).Path("/studentbook").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPost).Path("/studentbook").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPut).Path("/studentbook").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodDelete).Path("/studentbook").HandlerFunc(handler.StudentPostHandler)

	fmt.Println("server listening on", 8080)
	log.Fatal(http.ListenAndServe(":8080", router))
}
