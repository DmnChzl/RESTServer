package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrdoomy/restserver/controllers"
	"github.com/mrdoomy/restserver/utils"

	"github.com/gorilla/mux"
)

// PORT : Listening Port
const PORT = "4321"

func init() {
	utils.LoadFile("books.json")
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	api.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	api.HandleFunc("/book", controllers.AddBook).Methods("POST")
	api.HandleFunc("/book/{id}", controllers.ModifyBook).Methods("PUT")
	api.HandleFunc("/books", controllers.RemoveAllBooks).Methods("DELETE")
	api.HandleFunc("/book/{id}", controllers.RemoveBook).Methods("DELETE")
	fmt.Println("Starting Server On Port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
