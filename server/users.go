package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"sort"
	"unicode"
	"unicode/utf8"
)

type MessagedUsers struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func updateOnlineUserList(removeUser string) {
	delete(users, removeUser)
}

func removeByIndex(slice []string, index int) []string {

	return append(slice[:index], slice[index+1:]...)
}

func findIndex(list []string, el string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == el {
			return i
		}
	}
	return -1
}

func getMessagedUsernames(user string) []string {
	var messagedUsers []string

	db, err := sql.Open("sqlite3", "./forum.db")

	rows, err := db.Query("SELECT sender, receiver from messages where sender=? OR receiver=? ORDER BY id DESC", user, user)
	if err != nil {
		fmt.Println("error getting messaged users", err)
	}
	defer db.Close()

	for rows.Next() {
		message := MessagedUsers{}
		err = rows.Scan(&message.Sender, &message.Receiver)
		if err != nil {
			log.Println(err)
		}

		if message.Sender != user {
			if !contains(messagedUsers, message.Sender) {
				messagedUsers = append(messagedUsers, message.Sender)
			}
		} else {
			if !contains(messagedUsers, message.Receiver) {
				messagedUsers = append(messagedUsers, message.Receiver)
			}
		}
	}
	return messagedUsers
}

func lessLower(sa, sb string) bool {
	for {
		rb, nb := utf8.DecodeRuneInString(sb)
		if nb == 0 {
			// The number of runes in sa is greater than or
			// equal to the number of runes in sb. It follows
			// that sa is not less than sb.
			return false
		}

		ra, na := utf8.DecodeRuneInString(sa)
		if na == 0 {
			// The number of runes in sa is less than the
			// number of runes in sb. It follows that sa
			// is less than sb.
			return true
		}

		rb = unicode.ToLower(rb)
		ra = unicode.ToLower(ra)

		if ra != rb {
			return ra < rb
		}

		// Trim rune from the beginning of each string.
		sa = sa[na:]
		sb = sb[nb:]
	}
}

func usersOnlineList(w http.ResponseWriter, r *http.Request) {

	w = ConfigHeader(w)

	var user string
	var onlineUsers []string
	var onlineUserList []string
	var onlineWithoutMessages []string

	sessionID, errCookie := r.Cookie("session")
	if errCookie != nil {
		fmt.Println("Browser does not have a cookie!", errCookie)
		RespondWithError(w, "errCookie", 200)
		return
	}

	db, _ := sql.Open("sqlite3", "./forum.db")
	row := db.QueryRow("select username from session where session_id = ?", sessionID.Value)
	row.Scan(&user)

	onlineUsers = UserNameSlice
	sort.Slice(onlineUsers, func(i, j int) bool { return lessLower(onlineUsers[i], onlineUsers[j]) })

	var messagedUsers = getMessagedUsernames(user)

	if len(onlineUsers) > 0 {
		for _, v := range messagedUsers {
			if contains(onlineUsers, v) {
				onlineUserList = append(onlineUserList, v)
			}
		}

		for _, v := range onlineUsers {
			if !contains(messagedUsers, v) {
				onlineWithoutMessages = append(onlineWithoutMessages, v)
			}
		}
	}

	// adding rest of the users
	if len(onlineWithoutMessages) > 0 {
		for _, v := range onlineWithoutMessages {
			onlineUserList = append(onlineUserList, v)
		}
	}

	fmt.Println("OnlineUserList ", onlineUserList)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(onlineUserList); err != nil {
		panic(err)
	}
}

func getOnlineUsersWS(user string) []string {

	var onlineUsers []string
	var onlineWithoutMessages []string
	var onlineUserList []string

	onlineUsers = UserNameSlice
	sort.Slice(onlineUsers, func(i, j int) bool { return lessLower(onlineUsers[i], onlineUsers[j]) })

	var messagedUsers = getMessagedUsernames(user)

	if len(onlineUsers) > 0 {
		for _, v := range messagedUsers {
			if contains(onlineUsers, v) {
				onlineUserList = append(onlineUserList, v)
			}
		}

		for _, v := range onlineUsers {
			if !contains(messagedUsers, v) {
				onlineWithoutMessages = append(onlineWithoutMessages, v)
			}
		}
	}

	// adding rest of the users
	if len(onlineWithoutMessages) > 0 {
		for _, v := range onlineWithoutMessages {
			onlineUserList = append(onlineUserList, v)
		}
	}
	return onlineUserList
}
