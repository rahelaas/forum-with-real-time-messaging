package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var users = make(map[string]*websocket.Conn)

// receiving a message from frontend
type ChatMessage struct {
	Type           string   `json:"type"`
	Sender         string   `json:"sender"`
	Receiver       string   `json:"receiver"`
	Content        string   `json:"body"`
	Date           string   `json:"date"`
	OnlineUserList []string `json:"onlineUsers"`
}

var cookies *http.Cookie

var UserNameSlice = make([]string, 0, len(users))

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func RemoveUser(s []string, r string) []string {
	if contains(s, r) {

		for i, name := range s {
			if name == r {
				if i == 0 {
					s = s[i+1:]
				} else if i > 0 {

					s = append(s[:i], s[i+1:]...)
				}
			}
		}
		return s
	}
	return s
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	w = ConfigHeader(w)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Get user cookie
	var err3 error
	cookie, err3 := r.Cookie("session")
	if err3 != nil {
		log.Println(err3)
	}
	cookies = cookie

	db, _ := sql.Open("sqlite3", "./forum.db")
	userName := ""
	row := db.QueryRow("select username from session WHERE session_id= ?", cookie.Value)
	row.Scan(&userName)

	for userN := range users {
		if userN == userName || userN == "" {
			delete(users, userName)
		}
	}

	if _, ok := users[userName]; !ok {
		users[userName] = ws
		fmt.Println("Websocket connections: ", users)

	}

	UserNameSlice = make([]string, 0, len(users))
	for nameOfUser := range users {
		UserNameSlice = append(UserNameSlice, nameOfUser)
	}
	fmt.Println("Ongoing connections: ", len(users))

	go Writer(ws)
}

// define a writer which will send
// new messages to our WebSocket endpoint
func Writer(ws *websocket.Conn) {

	var usersForLoop = make(map[string]*websocket.Conn)
	for {
		code, content, err := ws.ReadMessage()
		if err != nil {
			log.Println("err on reading message in writer", err)
			log.Println("err on reading content in writer", content)
			log.Println("err on code message in writer", code)

			return
		}

		var chatMessage ChatMessage
		err7 := json.Unmarshal(content, &chatMessage)
		if err7 != nil {
			fmt.Println("error unMarshaling:", err7)

		}
		if chatMessage.Type == "chatMessage" {

			InsertMessageIntoDB(chatMessage.Sender, chatMessage.Receiver, chatMessage.Content)
			usersForLoop = make(map[string]*websocket.Conn)
			for user := range users {
				if user == chatMessage.Sender || user == chatMessage.Receiver {
					usersForLoop[user] = users[user]
				}
			}
		} else if chatMessage.Type == "userLeft" {
			for user := range users {
				if user == chatMessage.Sender {
					delete(users, user)
					ws.Close()
				}
				usersForLoop = users

			}
		} else if chatMessage.Type == "userJoined" {
			usersForLoop = users
		} else if chatMessage.Type == "startTyping" {
			usersForLoop = make(map[string]*websocket.Conn)
			for user := range users {
				if user == chatMessage.Receiver {
					usersForLoop[user] = users[user]
				}
			}
		} else if chatMessage.Type == "stopTyping" {
			usersForLoop = make(map[string]*websocket.Conn)
			for user := range users {
				if user == chatMessage.Receiver {
					usersForLoop[user] = users[user]
				}
			}
		}

		for user := range usersForLoop {
			if chatMessage.Type == "userJoined" || chatMessage.Type == "chatMessage" {
				onlineUsers := getOnlineUsersWS(user)
				chatMessage.OnlineUserList = onlineUsers
				content, _ = json.Marshal(chatMessage)
			}

			w, err := usersForLoop[user].NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			// log.Println("sending message ", string(content))
			_, err = w.Write(content)
			if err != nil {
				log.Println("Error: ", err)
				return
			}

			if err := w.Close(); err != nil {
				log.Println("Error: ", err)
				return
			}
		}
	}
}
