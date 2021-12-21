package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/dmnchzl/restserver/dao"
	"github.com/dmnchzl/restserver/models"
)

// LoadFile : Populates DataBase With Some Data (Optional)
func LoadFile(filename string) {
	var books []models.Book

	count := len(dao.ReadAllBooks())

	// Load Values From JSON File To Model
	if count == 0 {
		byteValues, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(byteValues, &books)

		dao.CreateManyBooks(books)
	}
}
