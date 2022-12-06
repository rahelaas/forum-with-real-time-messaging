package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// data we receive
type CommentDataReceived struct {
	Username string `json:"username"`
	Content  string `json:"body"`
	PostID   int    `json:"post_id"`
}

func insertComment(w http.ResponseWriter, re *http.Request) {

	w = ConfigHeader(w)

	decoder := json.NewDecoder(re.Body)

	var commentDataReceived CommentDataReceived
	var msg message
	decoder.Decode(&commentDataReceived)

	db, _ := sql.Open("sqlite3", "./forum.db")
	stmt, _ := db.Prepare("INSERT INTO comments(username, body, post_id, date) values(?,?,?,DATETIME('now', 'localtime'))")
	stmt.Exec(commentDataReceived.Username, commentDataReceived.Content, commentDataReceived.PostID)
	fmt.Println("Inserting Comment into Database")
	msg.Msg = "Comment successfully created"
	msg.Check = true
	defer db.Close()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}
