package dao

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/mrdoomy/restserver/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// URI : MongoDB URI
const URI = "mongodb://localhost:27017"

// DB : DataBase Name
const DB = "book_store"

// COLL : Collection Name
const COLL = "books"

// var db *mongo.Database
var collection *mongo.Collection

// Establish New Connection To DataBase
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(DB)
	collection = db.Collection(COLL)
}

// CreateManyBooks : Create Many New Books Into DataBase From 'list' (CRUD)
func CreateManyBooks(books []models.Book) {
	var allBooks []interface{}
	for _, book := range books {
		allBooks = append(allBooks, book)
	}

	_, err := collection.InsertMany(context.Background(), allBooks)
	if err != nil {
		fmt.Println("[CreateManyBooks]: Failed To InsertMany()")
		// log.Fatal(err)
	}
}

// CreateBook : Create New Book Into DataBase From 'object' (CRUD)
func CreateBook(book models.Book) {
	_, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		fmt.Println("[CreateBook]: Failed To InsertOne()")
		// log.Fatal(err)
	}
}

// ReadAllBooks : Read All Book From DataBase (CRUD)
func ReadAllBooks() (result []*models.Book) {
	cur, err := collection.Find(context.Background(), bson.D{}, nil)
	if err != nil {
		fmt.Println("[ReadAllBooks]: Failed To Find()")
		log.Fatal(err)
	}

	// ForEach
	for cur.Next(context.Background()) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return
}

// ReadBook : Read Existing Book From DataBase With 'isbn' (CRUD)
func ReadBook(isbn string) (result models.Book) {
	book := bson.D{
		{Key: "isbn", Value: isbn},
	}
	err := collection.FindOne(context.Background(), book).Decode(&result)
	if err != nil {
		fmt.Println("[ReadBook]: Failed To FindOne()")
		// log.Fatal(err)
	}
	return
}

// UpdateBook : Update Existing Book From DataBase With 'object' (CRUD)
func UpdateBook(book interface{}, isbn string) (status models.Status) {
	_, err := collection.UpdateOne(
		context.Background(),
		bson.D{
			{Key: "isbn", Value: isbn},
		},
		bson.D{
			{Key: "$set", Value: book},
		})
	if err != nil {
		status.Result = "Failed To UpdateOne..."
		// log.Fatal(err)
	} else {
		status.Result = isbn + " Updated."
	}
	return
}

// DeleteAllBooks : Delete All Books From DataBase (CRUD)
func DeleteAllBooks() (status models.Status) {
	res, err := collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		status.Result = "Failed To DeleteMany..."
		// log.Fatal(err)
	} else {
		status.Result = strconv.FormatInt(res.DeletedCount, 10) + " Deleted."
	}
	return
}

// DeleteBook : Delete Existing Book From DataBase With 'isbn' (CRUD)
func DeleteBook(isbn string) (status models.Status) {
	_, err := collection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "isbn", Value: isbn},
		},
		nil)
	if err != nil {
		status.Result = "Failed To DeleteOne..."
		// log.Fatal(err)
	} else {
		status.Result = isbn + " Deleted."
	}
	return
}
