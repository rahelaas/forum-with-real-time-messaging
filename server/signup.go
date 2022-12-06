package main

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type userData struct {
	Username  string `json:"username"`
	Age       int    `json:"year"`
	Gender    string `json:"gender"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type message struct {
	Msg   string `json:"msg"`
	Check bool   `json:"check"`
}

func usernameExists(username string) bool {
	db, _ := sql.Open("sqlite3", "./forum.db")
	row := db.QueryRow("select username from user_account where username= ?", username)

	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

func emailExists(email string) bool {
	db, _ := sql.Open("sqlite3", "./forum.db")
	row := db.QueryRow("select email from user_account where email= ?", email)

	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

func signup(w http.ResponseWriter, request *http.Request) {

	w = ConfigHeader(w)

	decoder := json.NewDecoder(request.Body)

	var signUpData userData
	var msg message

	decoder.Decode(&signUpData)

	db, _ := sql.Open("sqlite3", "./forum.db")

	if usernameExists(signUpData.Username) && emailExists(signUpData.Email) {
		msg.Msg = "Username and email are already taken."
		msg.Check = false
	} else if emailExists(signUpData.Email) {
		msg.Msg = "Email is taken."
		msg.Check = false
	} else if usernameExists(signUpData.Username) {
		msg.Msg = "Username is taken."
		msg.Check = false

	} else if !emailExists(signUpData.Email) && !usernameExists(signUpData.Username) {

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(signUpData.Password), bcrypt.MinCost)

		stmt, _ := db.Prepare("INSERT INTO user_account(username, email, password, firstname, lastname, gender, yearOfBirth) values(?,?,?,?,?,?,?)")
		stmt.Exec(signUpData.Username, signUpData.Email, hashedPassword, signUpData.Firstname, signUpData.Lastname, signUpData.Gender, signUpData.Age)
		msg.Msg = "User successfully created"
		msg.Check = true

	}

	defer db.Close()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}

}
