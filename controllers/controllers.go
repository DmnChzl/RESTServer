package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dmnchzl/restserver/dao"
	"github.com/dmnchzl/restserver/models"

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
	payload := dao.UpdateBook(book, params["id"])
	json.NewEncoder(w).Encode(payload)
}

// RemoveAllBooks : Remove All Books (EndPoint)
func RemoveAllBooks(w http.ResponseWriter, r *http.Request) {
	payload := dao.DeleteAllBooks()
	json.NewEncoder(w).Encode(payload)
}

// RemoveBook : Remove Existing Book (EndPoint)
func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload := dao.DeleteBook(params["id"])
	json.NewEncoder(w).Encode(payload)
}
