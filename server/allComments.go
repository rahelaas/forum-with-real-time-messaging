package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// this gets all Posts from SQL, JSON it and sends info to FrontEnd
type CommentsDataSend struct {
	Username string `json:"username"`
	Content  string `json:"body"`
	Date     string `json:"date"`
}

type CommentsDataReceived struct {
	PostID int `json:"id"`
}

func allComments(w http.ResponseWriter, r *http.Request) {

	w = ConfigHeader(w)

	decoder := json.NewDecoder(r.Body)

	var commentsDataReceived CommentsDataReceived

	//decoded received data
	decoder.Decode(&commentsDataReceived)

	var comments []CommentsDataSend
	db, _ := sql.Open("sqlite3", "./forum.db")

	w.Header().Set("Content-Type", "application/json")

	rows, _ := db.Query("select username, body, date from comments where post_id = ?", commentsDataReceived.PostID)

	for rows.Next() {
		var username string
		var body string
		var date string

		rows.Scan(&username, &body, &date)
		comments = append(comments, CommentsDataSend{username, body, date})
	}

	commentsBytes, _ := json.Marshal(&comments)

	w.WriteHeader(http.StatusOK)

	w.Write(commentsBytes)
	db.Close()
}
