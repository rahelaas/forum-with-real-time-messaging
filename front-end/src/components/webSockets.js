import store from "@/store";
//import axios from "axios";

export {connectToWebSocket, closeWebSocketConn, connection, receivedMessage};
let receivedMessage = {}
let connection = null;
const webSocketAddress = "ws://localhost:9200/ws";

function connectToWebSocket() {
    connection = new WebSocket(webSocketAddress);
    console.log("OPEN NEW WEBSOCKET")
    // 0 opening, 1 open, 2 closing, 3 closed
    console.log("Websocket state", connection.readyState)

    connection.onopen = () => {
        let messageBackend = {type: "userJoined", sender: store.state.user.username}
        connection.send(JSON.stringify(messageBackend))

    }

    connection.onmessage = (event) => {
       // console.log("User received from webSockets.js ", event.data)
        let messageParsed = JSON.parse(event.data)
        let users = []
         users = store.state.usersOnline
        let messages = store.state.tenMessages

        if (messageParsed.type === "userJoined") {
            if (!users.includes(messageParsed.sender)) {
                if (messageParsed.sender === store.state.user.username) {
                    users.push(messageParsed.sender)

                } else {
                    if (messageParsed.onlineUsers !== null) {
                        users = []
                        users = messageParsed.onlineUsers
                    }
                }
            }
        } else if (messageParsed.type === "userLeft") {
            if (users.includes(messageParsed.sender)) {
                const index = users.indexOf(messageParsed.sender);
                users.splice(index, 1)
            }

        } else if (messageParsed.type === "startTyping") {
            if (messageParsed.receiver === store.state.user.username) {
                store.state.typer = messageParsed.sender
                store.state.isTyping = true
            }
        } else if (messageParsed.type === "stopTyping") {
            store.commit('resetIsTyping')
            store.commit("clearTyper")
        } else if (messageParsed.type === "chatMessage") {
            if (!document.getElementById("chattingWith")) {
                store.commit('clearIsChattingWith')
            }

            if (messageParsed.onlineUsers !== null) {
                users = []
                users = messageParsed.onlineUsers
            }
            if (typeof messages === "string") {
                messages = []
            }

            if (store.state.user.isChattingWith === messageParsed.sender || store.state.user.username === messageParsed.sender) {
                let messageAmount = store.state.messagesLoaded
                let sum = messageAmount + 1
                store.commit('changeMessagesLoaded', sum)
                messages.push(messageParsed)
                receivedMessage = {}
                receivedMessage = messageParsed

                setTimeout(() => {
                    scrollToBottom()
                }, 50)
            }
            if (window.location.pathname !== "/messages" || (window.location.pathname === "/messages" && messageParsed.sender !== store.state.user.isChattingWith && messageParsed.sender !== store.state.user.username ) || ( store.state.user.username !== messageParsed.sender && document.getElementById("chattingWith").innerHTML !== messageParsed.sender)) {
                alert(`You have a new message from ${messageParsed.sender}`)
            }
        }
        store.commit('changeTenMessages', messages)
        store.commit('changeUserOnlineList', users)
    }



    /*   connection.onclose = (event) => {
          console.log("websocket closing", event)

          let messageBackend = {type: "userLeft", sender: store.state.user.username}
          console.log("this we are sending", messageBackend)
          connection.send(JSON.stringify(messageBackend))
    }*/

}
function scrollToBottom(){
    let element = document.getElementById('messageArea')
    element.scrollIntoView({behavior: "smooth", block: "end", inline: "nearest"});
}


function closeWebSocketConn() {
    connection.close();
}
