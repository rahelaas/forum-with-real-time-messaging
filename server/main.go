package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func setupRoutes() {

	mux.HandleFunc("/ws", webSocketHandler)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/", login)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/post", Auth(insertPost))
	mux.HandleFunc("/allPosts", Auth(allPosts))
	mux.HandleFunc("/deletePost", deletePost)
	mux.HandleFunc("/comment", insertComment)
	mux.HandleFunc("/allComments", allComments)
	mux.HandleFunc("/singlePost", Auth(singlePost))
	mux.HandleFunc("/messages", getMessages)
	mux.HandleFunc("/usersOnlineList", Auth(usersOnlineList))
}

var port = 9200
var mux = http.NewServeMux()

func main() {

	createDatabase()
	setupRoutes()
	fmt.Println("Server is at your service on port: ", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), mux))

}
