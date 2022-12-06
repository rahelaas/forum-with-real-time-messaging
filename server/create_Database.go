package main

import (
	"database/sql"
	"fmt"
	"log"
	//"time"

	_ "github.com/mattn/go-sqlite3"
)

func createDatabase() {

	log.Println("Creating sqlite-database.db...")

	dbFile, err := sql.Open("sqlite3", "./forum.db") // Open SQLite file
	_, _ = dbFile.Exec("drop table if exists session")

	if err != nil {
		fmt.Println("This one")
		log.Fatal(err.Error())
	}

	createUserAcc(dbFile)
	createSession(dbFile)
	createPostTable(dbFile)
	createCommentTable(dbFile)
	createMessages(dbFile)

	defer dbFile.Close() // Defer Closing the database
}

func createUserAcc(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
    CREATE TABLE IF NOT EXISTS user_account  (
        username varchar(255) not null PRIMARY KEY,
        firstName varchar(255) not null,
        lastName varchar(255) not null,
        email varchar(255) not null,
        gender varchar(255) not null,
        password varchar(255) not null,
		yearOfBirth varchar(255) not null
        )
        `)
	statement.Exec()

	stmt, err := dbConn.Prepare(`INSERT INTO user_account(username, firstName, lastName, email, gender, password, yearOfBirth) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// password: 12345
	stmt.Exec("Dave", "David", "Hill", "dave@david.com", "Male", "$2a$04$VpiiTJja2U4z0dbE1BZ6c.gr/6a822j5gMYfF3lMnNcl5jhSi4aMa", "1979")
	stmt.Exec("Pete", "Peter", "Pound", "pete@peter.com", "Male", "$2a$04$VpiiTJja2U4z0dbE1BZ6c.gr/6a822j5gMYfF3lMnNcl5jhSi4aMa", "1985")
	stmt.Exec("Linda", "Linda", "Joy", "linda@linda.com", "Female", "$2a$04$VpiiTJja2U4z0dbE1BZ6c.gr/6a822j5gMYfF3lMnNcl5jhSi4aMa", "2018")
	stmt.Exec("Juri", "Juri", "Ratas", "juri@ratas.com", "Male", "$2a$04$VpiiTJja2U4z0dbE1BZ6c.gr/6a822j5gMYfF3lMnNcl5jhSi4aMa", "1969")
}

func createSession(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
	CREATE TABLE IF NOT EXISTS session  (
		session_id  varchar(255) not null PRIMARY KEY,
		username varchar(255) not null,
		max_age datetime not null,
	   FOREIGN KEY (username) REFERENCES user_account(username)
	 )		
		`)
	statement.Exec()
}

func createPostTable(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
    CREATE TABLE IF NOT EXISTS posts  (
        id INTEGER not null PRIMARY KEY AUTOINCREMENT,
        username varchar(255) not null,
        title varchar(255) not null,
        body varchar not null,
        category varchar(255) not null,
        date varchar(255) not null
        )
        `)
	statement.Exec()

	stmt, err := dbConn.Prepare(`INSERT INTO posts(id, username, title, body, category, date) VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec("1", "Juri", "Got bones?", "Looking for bones. Any good places in town?", "🐕 Dogs", "2022-11-11 20:43:10")
	stmt.Exec("2", "Linda", "Looking for a hiking companion", "Hello! I am going to hike the Pacific Crest Trail (PCT). If there is somebody with the same plan, write to me and let's join forces!", "🗺 Travel", "2022-11-11 17:19:55")
	stmt.Exec("3", "Linda", "Which car is the best?", "Hey guys! I am looking for a campervan!", "🚗 Cars", "2022-11-20 18:19:55")
	stmt.Exec("4", "Juri", "I want 4x4!", "Snow came, where can I change tyres?", "🚗 Cars", "2022-11-21 18:19:55")
}

func createCommentTable(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
	CREATE TABLE IF NOT EXISTS comments  (
		id INTEGER not null PRIMARY KEY AUTOINCREMENT,
		body  TEXT not null,
	   post_id INTEGER not null,
	   username varchar(255) not null,
	   date varchar(255) not null,
       FOREIGN KEY (post_id)
       REFERENCES posts (id)
       ON DELETE CASCADE,
	   FOREIGN KEY (username) REFERENCES user_account (username) ON DELETE CASCADE
	 )
	 `)

	statement.Exec()
	stmt, err := dbConn.Prepare(`INSERT INTO comments(id, body, post_id, username, date) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec("1", "Have you looked at the central market?", "1", "Dave", "2022-11-13 18:33:10")
	stmt.Exec("2", "I've always wanted to go but I just had a leg injury, so I won't be able to join you.", "2", "Juri", "2022-11-16 10:30:25")

}

func createMessages(dbConn *sql.DB) {
	statement, _ := dbConn.Prepare(`
	CREATE TABLE IF NOT EXISTS messages  (
		id  integer not null PRIMARY KEY AUTOINCREMENT,
	 	body  text not null,
		date  varchar(255) not null,
	   	sender varchar(255) not null,
	  	receiver varchar(255) not null
	 )
	 `)
	statement.Exec()
}
