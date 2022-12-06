package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// data we receive and send
type MessageDataReceived struct {
	Type     string `json:"type"`
	Id       int    `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"body"`
	Date     string `json:"date"`
}

func InsertMessageIntoDB(sender string, receiver string, content string) {
	db, err := sql.Open("sqlite3", "./forum.db")

	stmt, _ := db.Prepare("INSERT INTO messages(sender, receiver, body, date) values(?,?,?,DATETIME('now', 'localtime'))")
	if err != nil {
		fmt.Println("error inserting message to DB", err)
	}
	defer db.Close()
	stmt.Exec(sender, receiver, content)
}

func getMessagesFromDB(sender string, receiver string, messagesLoaded int) []MessageDataReceived {
	messages := []MessageDataReceived{}

	db, _ := sql.Open("sqlite3", "./forum.db")

	query := `SELECT * FROM (
SELECT * FROM messages WHERE (sender=? AND receiver=?) OR (sender=? AND receiver=?)
                                        ORDER BY id DESC
                                		LIMIT 10 OFFSET ?
                                	)
                                	ORDER BY id ASC
                                	`

	rows, err := db.Query(query, sender, receiver, receiver, sender, messagesLoaded)
	if err != nil {
		fmt.Println("error ", err)
	}
	defer db.Close()

	for rows.Next() {
		message := MessageDataReceived{}
		err = rows.Scan(&message.Id, &message.Content, &message.Date, &message.Sender, &message.Receiver)
		if err != nil {
			fmt.Println(err)
		}
		messages = append(messages, message)
	}

	return messages
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	w = ConfigHeader(w)
	if r.Method == "POST" {
		// Get user cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Println(err)
		}
		db, _ := sql.Open("sqlite3", "./forum.db")
		username := ""
		row := db.QueryRow("select username from session WHERE session_id= ?", cookie.Value)
		row.Scan(&username)

		type fromFrontEnd struct {
			ChattingWith   string
			MessagesLoaded int
		}
		data := fromFrontEnd{}
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&data)

		messages := getMessagesFromDB(username, data.ChattingWith, data.MessagesLoaded)

		if err := json.NewEncoder(w).Encode(messages); err != nil {
			log.Println("Sending message error ", err)
		}
	}
}
