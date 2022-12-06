package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gofrs/uuid"
	//"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// data we receive
type loginDataRes struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}

// data we send back to client
type loginDataSend struct {
	FirstName string `json:"firstName"`
	Username  string `json:"username"`
	Msg       string `json:"msg"`
	Check     bool   `json:"check"`
}

func credentialCheck(credential string) bool {

	db, _ := sql.Open("sqlite3", "./forum.db")
	row := db.QueryRow("select username from user_account where username= ?", credential)
	row2 := db.QueryRow("select username from user_account where email = ?", credential)
	temp := ""
	temp2 := ""
	row.Scan(&temp)
	row2.Scan(&temp2)

	if temp != "" || temp2 != "" {
		return true
	}

	return false
}

func CheckPasswordHash(password, credential string) error {
	db, _ := sql.Open("sqlite3", "./forum.db")
	row := db.QueryRow("select password from user_account where username= ?", credential)

	temp := ""
	row.Scan(&temp)
	if temp == "" {
		row := db.QueryRow("select password from user_account where email= ?", credential)
		row.Scan(&temp)
	}
	err := bcrypt.CompareHashAndPassword([]byte(temp), []byte(password))
	return err
}

func login(w http.ResponseWriter, req *http.Request) {
	w = ConfigHeader(w)

	decoder := json.NewDecoder(req.Body)

	var loginDataRes loginDataRes
	var loginDataSend loginDataSend

	//decoded received data
	decoder.Decode(&loginDataRes)
	db, _ := sql.Open("sqlite3", "./forum.db")

	if !credentialCheck(loginDataRes.Credential) {
		loginDataSend.Msg = "Username or Email is wrong"
		loginDataSend.Check = false
	} else {
		err := CheckPasswordHash(loginDataRes.Password, loginDataRes.Credential)
		if err != nil {
			loginDataSend.Msg = "Password is wrong"
			loginDataSend.Check = false
		} else {
			row1 := db.QueryRow("select username from user_account where username= ?", loginDataRes.Credential)
			row2 := db.QueryRow("select username from user_account where email = ?", loginDataRes.Credential)
			row3 := db.QueryRow("select firstName from user_account where email = ?", loginDataRes.Credential)
			row4 := db.QueryRow("select firstName from user_account where username = ?", loginDataRes.Credential)
			temp := ""
			temp2 := ""
			temp3 := ""
			temp4 := ""
			row1.Scan(&temp)
			row2.Scan(&temp2)
			row3.Scan(&temp3)
			row4.Scan(&temp4)

			if temp != "" {
				loginDataSend.Username = temp
				loginDataSend.FirstName = temp4

			}
			if temp2 != "" {
				loginDataSend.Username = temp2
				loginDataSend.FirstName = temp3

			}

			loginDataSend.Msg = "Logged in successfully"
			loginDataSend.Check = true
			//sID := uuid.NewV4()
			sID, er := uuid.NewV4()
			if er != nil {
				log.Fatalf("failed to generate UUID: %v", er)
			}
			sessionID := sID.String()
			maxAge := 60 * 30

			cookie := http.Cookie{
				Name:     "session",
				Value:    sessionID,
				Secure:   false,
				Path:     "/",
				HttpOnly: false,
				MaxAge:   maxAge,
			}
			http.SetCookie(w, &cookie)

			// adds user to database and updates session
			updateSession(db, sID.String(), loginDataSend.Username)
		}
	}

	defer db.Close()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginDataSend); err != nil {
		panic(err)
	}
}

func updateSession(db *sql.DB, session_id string, username string) {
	var i int

	rows, _ := db.Query("SELECT session_id, username FROM session WHERE username = ?", (username))
	//checkErr(err)

	for rows.Next() {
		i++
	}

	if i > 0 {
		stmt, _ := db.Prepare("DELETE FROM session WHERE username = ?")
		//checkErr(err)
		stmt.Exec(username)
	}

	stmt, _ := db.Prepare(`INSERT INTO session(session_id, username, max_age) values(?,?, datetime("now", "+30 minutes"))`)
	//checkErr(err)

	stmt.Exec(session_id, username)
}
