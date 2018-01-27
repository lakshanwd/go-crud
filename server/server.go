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

	router.Methods(http.MethodGet).Path("/student").HandlerFunc(handler.StudentGetHandler)
	router.Methods(http.MethodPost).Path("/student").HandlerFunc(handler.StudentPostHandler)
	router.Methods(http.MethodPut).Path("/student").HandlerFunc(handler.StudentPutHandler)
	router.Methods(http.MethodDelete).Path("/student/{id:[0-9]+}").HandlerFunc(handler.StudentDeleteHandler)

	router.Methods(http.MethodGet).Path("/book").HandlerFunc(handler.BookGetHandler)
	router.Methods(http.MethodPost).Path("/book").HandlerFunc(handler.BookPostHandler)
	router.Methods(http.MethodPut).Path("/book").HandlerFunc(handler.BookPutHandler)
	router.Methods(http.MethodDelete).Path("/book/{id:[0-9]+}").HandlerFunc(handler.BookDeleteHandler)

	router.Methods(http.MethodGet).Path("/studentbook").HandlerFunc(handler.StudentBookGetHandler)
	router.Methods(http.MethodPost).Path("/studentbook").HandlerFunc(handler.StudentBookPostHandler)
	router.Methods(http.MethodPut).Path("/studentbook").HandlerFunc(handler.StudentBookPutHandler)
	router.Methods(http.MethodDelete).Path("/studentbook/{id:[0-9]+}").HandlerFunc(handler.StudentBookDeleteHandler)

	fmt.Println("server listening on", 8080)
	log.Fatal(http.ListenAndServe(":8080", router))
}
