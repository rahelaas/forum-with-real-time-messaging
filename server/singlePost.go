package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// this gets all Posts from SQL, JSON it and sends info to FrontEnd
type SinglePostDataSend struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"body"`
	Category string `json:"category"`
	PostID   int    `json:"id"`
	Date     string `json:"date"`
}

type SinglePostDataReceived struct {
	PostID int `json:"id"`
}

func singlePost(w http.ResponseWriter, r *http.Request) {

	w = ConfigHeader(w)

	decoder := json.NewDecoder(r.Body)

	var singlePostDataReceived SinglePostDataReceived
	var singlePostDataSend SinglePostDataSend

	//decoded received data
	decoder.Decode(&singlePostDataReceived)

	db, _ := sql.Open("sqlite3", "./forum.db")

	w.Header().Set("Content-Type", "application/json")

	row := db.QueryRow("select * from posts where id = ?", singlePostDataReceived.PostID)

	row.Scan(&singlePostDataSend.PostID, &singlePostDataSend.Username, &singlePostDataSend.Title, &singlePostDataSend.Content, &singlePostDataSend.Category, &singlePostDataSend.Date)

	postBytes, _ := json.Marshal(&singlePostDataSend)

	w.WriteHeader(http.StatusOK)

	w.Write(postBytes)
	db.Close()

}
