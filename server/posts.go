package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// data we receive
type PostDataReceived struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"body"`
	Category string `json:"category"`
}

type PostDataSend struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"body"`
	Category string `json:"category"`
	Date     string `json:"date"`
	PostID   int    `json:"id"`
}

// this gets all Posts from SQL, JSON it and sends info to FrontEnd
func allPosts(w http.ResponseWriter, r *http.Request) {

	w = ConfigHeader(w)

	var posts []PostDataSend
	db, err := sql.Open("sqlite3", "./forum.db")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select username, title, body, category, date, id from posts ORDER BY id DESC;")

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var username string
		var title string
		var body string
		var category string
		var date string
		var id int

		rows.Scan(&username, &title, &body, &category, &date, &id)
		posts = append(posts, PostDataSend{username, title, body, category, date, id})
	}
	postsBytes, err := json.Marshal(posts)
	if err != nil {
	}
	w.Write(postsBytes)
	db.Close()
}

func insertPost(w http.ResponseWriter, r *http.Request) {

	w = ConfigHeader(w)

	decoder := json.NewDecoder(r.Body)

	var postDataReceived PostDataReceived
	var msg message
	decoder.Decode(&postDataReceived)

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt, _ := db.Prepare("INSERT INTO posts(username, title, body, category, date) values(?,?,?,?,DATETIME('now', 'localtime'))")
	stmt.Exec(postDataReceived.Username, postDataReceived.Title, postDataReceived.Content, postDataReceived.Category)
	fmt.Println("Inserting Post into Database")
	msg.Msg = "Post successfully created"
	msg.Check = true
	defer db.Close()

	w.WriteHeader(http.StatusOK)
	if er := json.NewEncoder(w).Encode(msg); er != nil {
		fmt.Println("Error on sending message")
		panic(er)
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {

	w = ConfigHeader(w)

	decoder := json.NewDecoder(r.Body)
	var id int
	var msg message

	decoder.Decode(&id)
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("PRAGMA foreign_keys = ON")
	defer db.Close()

	res, err := db.Exec("DELETE FROM posts WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		msg.Msg = "Something went wrong with deleting the post"
		msg.Check = false
	} else {
		msg.Msg = "Post successfully deleted"
		msg.Check = true
	}
	defer db.Close()

	n, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The statement deleted %d row(s)\n", n)
	fmt.Println()
	if er := json.NewEncoder(w).Encode(msg); er != nil {
		fmt.Println("Error on sending message")
		panic(er)
	}
}
