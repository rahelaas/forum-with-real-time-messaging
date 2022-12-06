package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// data we receive
type LogoutDataRes struct {
	Username string `json:"usernameOffline"`
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w = ConfigHeader(w)

	decoder := json.NewDecoder(r.Body)
	var logoutDataRes LogoutDataRes

	decoder.Decode(&logoutDataRes)

	cookie := http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: false,
		MaxAge:   -99999,
	}
	http.SetCookie(w, &cookie)

	db, err1 := sql.Open("sqlite3", "./forum.db")

	if err1 != nil {
		log.Fatal(err1)
	}

	defer db.Close()

	res, err := db.Exec("DELETE FROM session WHERE username = ?", logoutDataRes.Username)

	if err != nil {
		log.Fatal(err)
	}

	n, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The statement deleted %d user(s) from session table", n)
	fmt.Println()

	UserNameSlice = RemoveUser(UserNameSlice, logoutDataRes.Username)

}
