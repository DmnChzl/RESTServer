package models

// Book As Model
type Book struct {
	Isbn      string   `json:"isbn"`
	Title     string   `json:"title"`
	SubTitle  string   `json:"subTitle" bson:"subTitle"`
	Volume    int      `json:"volume"`
	Date      string   `json:"date"`
	Authors   []Author `json:"authors"`
	Publisher string   `json:"publisher"`
	Category  string   `json:"category"`
	Chapters  []string `json:"chapters"`
	Pages     int      `json:"pages"`
	Price     float64  `json:"price"`
}

// Author As Model
type Author struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}
