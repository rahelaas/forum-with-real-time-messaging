package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ResponseMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"` // message itself
}

func RespondWithError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	err := ResponseMessage{Message: message, Type: "Error"}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}

type SessionInfo struct {
	UserName string
	MaxAge   time.Time
}

var sessionInfo SessionInfo

func CreateCookie(sessionID string, lifespan int) http.Cookie {
	return http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   lifespan,
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w = ConfigHeader(w)

		sessionID, errCookie := r.Cookie("session")
		if errCookie != nil {
			//fmt.Println("Browser does not have a cookie!", errCookie)
			RespondWithError(w, "errCookie", 200)
			return
		}

		db, _ := sql.Open("sqlite3", "./forum.db")
		row := db.QueryRow("select username, max_age from session where session_id = ?", sessionID.Value)
		row.Scan(&sessionInfo.UserName, &sessionInfo.MaxAge)

		if sessionInfo.MaxAge.After(time.Now()) {
			updateSession(db, sessionID.Value, sessionInfo.UserName)
			cookie := http.Cookie{
				Name:     "session",
				Value:    sessionID.Value,
				Path:     "/",
				HttpOnly: false,
				MaxAge:   60 * 30,
			}
			http.SetCookie(w, &cookie)
		} else {
			cookie := CreateCookie(" ", -1)
			http.SetCookie(w, &cookie)
			stmt, err := db.Exec("DELETE FROM session WHERE username = ?", sessionInfo.UserName)
			if err != nil {
				log.Fatal("doesnt work session cookie", err)
			}
			log.Println("some statement", stmt)
		}
		next(w, r)
	})
}
