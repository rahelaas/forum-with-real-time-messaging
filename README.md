# **Forum with Real-Time Messaging Showing Typing in Progress**
(original name Real-Time Forum Typing-in-Progress)

## **Objectives**
This project consists in creating a web forum that allows:

* Registration and Login
* Creation of posts
* Commenting posts
* Private Messages
* A Typing in progress engine

This forum has five different parts:


* SQLite, in which to store data
* Golang, in which the data and Websockets (Backend) are handled
* Javascript, in which the Frontend events and clients Websockets are handled
* HTML, in which to organize the elements of the page
* CSS, in which to stylize the elements of the page

The backend is written in [Golang](https://go.dev/) and the Frontend is written in [Vue](https://vuejs.org/) - a JavaScript framework. The task limitation to use JavaScript without a framework came after finishing this project.

The Go websockets are created with Gorilla websocket package.

## **How to run:**

You need to run both back-end and front-end servers in separate terminals.

### To run back-end server (_GO_) from the project folder:
```
cd server
go run .
```

### To run front-end server (_VUEJS_) from the project folder:
```
cd front-end
npm install
npm run serve
```
## When both servers are up & running, go to the following address in your browser:
```
http://localhost:8080
```
## Full task description

[Task description](https://github.com/01-edu/public/tree/master/subjects/real-time-forum/typing-in-progress)

## Authors:
```
kaspars & rahela
```