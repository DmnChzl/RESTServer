package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mrdoomy/restserver/dao"
	"github.com/mrdoomy/restserver/models"

	"github.com/gorilla/mux"
)

// GetAllBooks : Get All Books (EndPoint)
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	payload := dao.ReadAllBooks()
	json.NewEncoder(w).Encode(payload)
}

// GetBook : Get Existing Book (EndPoint)
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload := dao.ReadBook(params["id"])
	json.NewEncoder(w).Encode(payload)
}

// AddBook : Add New Book (EndPoint)
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	dao.CreateBook(book)
	json.NewEncoder(w).Encode(book)
}

// ModifyBook : Modify Existing Book (EndPoint)
func ModifyBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book interface{}
	_ = json.NewDecoder(r.Body).Decode(&book)
	dao.UpdateBook(book, params["id"])
}

// RemoveAllBooks : Remove All Books (EndPoint)
func RemoveAllBooks(w http.ResponseWriter, r *http.Request) {
	dao.DeleteAllBooks()
}

// RemoveBook : Remove Existing Book (EndPoint)
func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dao.DeleteBook(params["id"])
}
