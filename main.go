package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

type Post struct {
	ID     int   `json:"id,omitempty"`
	PostId string   `json:"postid"`
}

// Display a single data
//func GetPerson(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	for _, item := range people {
//		if item.ID == params["id"] {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&Person{})
//}

// create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	db, err := sql.Open("mysql", "tntest:EF.tn.t3sTdB@tcp(localhost)/test_db")

	if err != nil {
		fmt.Printf("Unable to open db: %s\n", err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// perform a db.Query insert
	query := fmt.Sprintf("INSERT INTO posts (post_id) VALUES ('%s')", post.PostId)
	fmt.Print(query)
	insert, err := db.Query(query)


	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	json.NewEncoder(w).Encode(post)
}

// Delete an item
//func DeletePerson(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	for index, item := range people {
//		if item.ID == params["id"] {
//			people = append(people[:index], people[index+1:]...)
//			break
//		}
//		json.NewEncoder(w).Encode(people)
//	}
//}

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", CreatePost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}