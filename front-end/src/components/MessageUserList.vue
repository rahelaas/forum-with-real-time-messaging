<template>
  <section class="parent">
  <section class="leftbox">
    <h3 class="users-headline-left">Online Users</h3>
  <div class="user-list-area">
  <div v-if=" userOnlineList.length > 0 ">
    <ul v-for="userOnline in onlineList"
        :key="userOnline"
        style="padding:0rem"
    >
      <li class="user-list-item">
        <div class="single-user"  @click="handleUser(userOnline)"><span class="logged-in">● </span>{{userOnline}} <div v-if="isTyping && userOnline === typer" class="typing">
          <div class="cp_typing_indicator">
            <span class="cp_ball cp_ball1"></span>
            <span class="cp_ball cp_ball2"></span>
            <span class="cp_ball cp_ball3"></span>
          </div>
        </div>
        </div>
      </li>
    </ul>
    </div>
    <div class="noUsersOnline" v-else>No online users</div>
  </div>
  </section>

  <section class="rightbox">
    <div v-if="message.receiver !== ''" class="users-headline-right">You are chatting with:&nbsp;
      <div id="chattingWith" class="chattingWith">{{message.receiver}}</div></div>
    <div v-else-if="userOnlineList.length > 0" class="users-headline-right">Pick a friend to chat with</div>
    <div v-else class="users-headline-right">Waiting for online users <div id="loading"></div></div>
  <section ref="chatArea" class="chat-area" v-on:scroll="doingThrottle">
    <div id="messageArea">
      <ul v-for="oneMessage in tenMessages"
          :key="oneMessage"
      >
        <div v-if="oneMessage.sender === currentUser" class="message" :class="{ 'message-out': oneMessage.body }"> {{oneMessage["body"] }}
          <br>
          <div class="chatReceiver">
            {{oneMessage.sender}} - Sent: {{oneMessage.date}}
          </div>
        </div>
        <div v-else class="message" :class="{ 'message-in': oneMessage.body }">
          {{oneMessage["body"] }}
          <br>
          <div class="chatSender">
            {{oneMessage.sender}} - Sent: {{oneMessage.date}}
          </div>
        </div>
      </ul>
    </div>
  </section>
  </section>

  <section class="chat-inputs">
    <form @submit.prevent>
      <my-input id="title" class="title" v-model.trim="currentMessage" type="text" placeholder="Type your message" autocomplete="off" required @keydown="onTyping($event)"/>
      <section class="typo">
        <!--if isTyping= true and message receiver has opened the chat window with message sender--->
        <div v-if="isTyping && chatWindowOpen && message.receiver === typer" class="typing"> {{ typer }} is typing
          <div class="cp_typing_indicator">
            <span class="cp_ball cp_ball1"></span>
            <span class="cp_ball cp_ball2"></span>
            <span class="cp_ball cp_ball3"></span>
          </div>
        </div>
      </section>
      <my-button class="btn" value="Send" @click.prevent="sendMessage()">Send</my-button>
    </form>
  </section>
  </section>

</template>

<script>
import axios from "axios";
import _, {debounce} from 'lodash';
axios.defaults.withCredentials = true;
import store from "@/store";
import {connection, receivedMessage} from "@/components/webSockets";

export default {
  name: "MessageUserList",
  data() {
    return {
      currentUser: store.state.user.username,
      currentMessage: "",
      message: {
        type: "",
        id: "",
        sender: store.state.user.username,
        receiver: "",
        body: "",
        date: "",
      },
      receivedChatMsg: receivedMessage,
      tenMessages: [],
      userOnlineList: [],
      chatWindowOpen: false,
    }
  },
  computed: {
    onlineList() {
      let onlineUsers = store.getters.getOnlineUsers
      return onlineUsers.filter((name) => {
        return this.message.sender != name
      })
    },
    isTyping() {
      return store.state.isTyping
    },
    typer() {
      return store.state.typer
    },
  },
  watch: {
    onlineList(newVal) {
      this.userOnlineList = newVal
    },
    tenMessages(now, then){
      if (now.length != then.length && now.length > 0){
        this.scrollUpdate = true
        this.scrollToMsg = then.length !=0 ? now.length - then.length + 5: now.length;
      }
    }
  },

  updated(){
    if(this.scrollUpdate){
      this.scrollToBottom( this.scrollToMsg)
      this.scrollUpdate = false
    }
  },


  methods: {
    //sends message to Websockets.js if user starts/stops typing
    onTyping(e) {
      if( e.key != "Enter") {
      this.message.type = "startTyping"
      connection.send(JSON.stringify(this.message))
      this.debounceStopTyping();
    }
    },
    debounceStopTyping: debounce(function () {
      this.message.type = "stopTyping"
      connection.send(JSON.stringify(this.message))
    }, 500),

    scrollToBottom(child){
      if (child !== null) {
        let element = document.querySelector(`#messageArea :nth-child(${child})`)
        if (element !== null) {
          element.scrollIntoView(false);
        }

      }
    },

    doingThrottle: _.throttle(function ()  {

      if (this.$refs.chatArea.scrollTop == 0) {
        this.getMessagesFromSQL()
      }
    }, 700, {trailing: true, leading: true}),

    handleUser(userOnline) {
      this.message.receiver = userOnline
      store.state.user.isChattingWith = userOnline
      store.commit('changeMessagesLoaded', 0)
      this.tenMessages = []
      this.getMessagesFromSQL()
      this.chatWindowOpen = true
    },

    getUserOnlineList() {
      axios({
            url: "http://localhost:9200/usersOnlineList",
            method: "get",
            headers: {'Content-Type': 'application/json'},
            withCredentials: true,
            origin: true,
            //  origin: "http://localhost:8080",
          },
      ).then((response) => {
        store.state.usersOnline = response.data
              this.userOnlineList = store.state.usersOnline
          },
      ).catch((error) => {
        console.log("Error from getting online users ", error)
      })
    },

    getMessagesFromSQL() {
      axios({
        method: "post",
        url: "http://localhost:9200/messages",
        data: {
          chattingWith: this.message.receiver,
          messagesLoaded: store.state.messagesLoaded,
        },
        headers: {'Content-Type': 'application/json'},
        withCredentials: true,
        origin: true,
        // origin: "http://localhost:8080",
      }).then((result) => {

        if (store.state.messagesLoaded === 0) {
          store.state.tenMessages = result.data
          this.tenMessages = store.state.tenMessages
          let messageAmount = store.state.messagesLoaded
          let sum = messageAmount + result.data.length
          store.commit('changeMessagesLoaded', sum)
        } else {
          let mm = result.data.concat(this.tenMessages)
          store.commit('changeTenMessages', mm)
          this.tenMessages = store.state.tenMessages
          let messageAmount = store.state.messagesLoaded
          let sum = messageAmount + result.data.length
          store.commit('changeMessagesLoaded', sum)
        }
      }).catch((error) => {
        if (error.response) {
          console.log(error.response);
        } else if (error.request) {
          console.log("network error");
        } else {
          console.log(error);
        }
      });
    },

    sendMessage() {
      if (this.message.receiver === "") {
        alert("Pick the message receiver")
      } else if (this.currentMessage !==""){
        const dt = new Date();
        const padL = (nr, chr = `0`) => `${nr}`.padStart(2, chr);
        let date = `${dt.getFullYear()}-${
          padL(dt.getMonth()+1)}-${
            padL(dt.getDate())} ${
            padL(dt.getHours())}:${
            padL(dt.getMinutes())}:${
            padL(dt.getSeconds())}`;

        this.message.date = date
        this.message.body = this.currentMessage,
            this.message.type = "chatMessage"

          // 0 opening, 1 open, 2 closing, 3 closed
        console.log("Websocket state before sending", connection.readyState)
        setTimeout( () => {
          connection.send(JSON.stringify(this.message))
          this.currentMessage = ""
        }, 50)
      }
    },
  },

  mounted() {
    this.getUserOnlineList()
   // this.onTyping()
    },
}
</script>


<style scoped>
.parent {
  display: flex;
  position:fixed;
  gap: 10px;
  justify-content: center;
  flex-wrap:wrap;
  min-width: 1200px;
  max-width: 1300px;
}
.rightbox {
  justify-content: center;
  align-items: center;
}
.leftbox {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.users-headline-left {
  width: 100%;
  margin-top: 5%;
  margin-bottom: 7%;
}
.logged-in {
  color: teal;
}
.users-headline-right {
  width: 100%;
  margin-top: 2%;
  margin-bottom: 3%;
  justify-content: center;
  align-items: center;
  display: flex;
}
.chattingWith {
  font-weight: bold;
  color: teal;
}
.user-list-area{
  align-items: center;
  text-align: center;
  justify-content: center;
  float:left; display:inline; width: 100%;
  top: 150%;
  background: white;
  height: 50vh;
  padding: 1em;
  overflow: auto;
  max-width:200px;
  min-width:150px;
  margin: 1%;
  box-shadow: 2px 2px 5px 2px rgba(0, 0, 0, 0.3)
}
.chat-area {
  background: white;
  height: 50vh;
  width: 100vh;
  padding: 1em;
  overflow: auto;
  max-width:850px;
  margin: 1%;
  box-shadow: 2px 2px 5px 2px rgba(0, 0, 0, 0.3);
}

.message {
  width: 50%;
  border-radius: 10px;
  padding: .5em;
  font-size: .8em;
}
.message-out {
  background: #407FFF;
  color: white;
  margin-left: 50%;
}
.message-in {
  background: #F1F0F0;
  color: black;
}
.chatSender {
  font-style: italic;
  color: teal;
  font-size: 10px;
}
.chatReceiver {
  font-style: italic;
  color: white;
  font-size: 10px;
}
.chat-inputs {
  width: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
form {
  left: 200px;
  width: 90vh;
}
.noUsersOnline {
  alignment: center;
}
.user-list-item {
  width: 100%;
  padding: 0.5rem 1rem;
  list-style: none;
  text-align: left;

}
.single-user{
  padding: 0%;
  width: fit-content;
  display: flex;
}
.single-user:hover{
  background-color: hsla(120,100%,25%,0.07);
}

/*SPINNER*/
#loading {

  display: inline-block;
  width: 60px;
  height: 60px;
  border: 2px solid rgb(0 188 0 / 20%);
  border-radius: 50%;
  border-top-color: teal;
  animation: spin 1s ease-in infinite;
  -webkit-animation: spin 1s ease-in-out infinite;
  position: absolute;
  top: 42%;
  left: 56%;
  box-sizing: border-box;
  translate: -50% -50%;
}
@keyframes spin {
  to {
    -webkit-transform: rotate(360deg);
  }
}
@-webkit-keyframes spin {
  to {
    -webkit-transform: rotate(360deg);
  }
}

/*typing */
.typo {
  padding: 5px;
  position: absolute;

}
.typing {
  font-style: italic;
  color: darkolivegreen;
}
.cp_typing_indicator {
  display: inline-block;
}

.cp_ball {
  display: inline-block;
  width: 7px;
  height: 7px;
  margin: 0 1px;
  border-radius: 50%;
  background-color: darkolivegreen;
  animation: ball_typing 1.2s linear infinite;
}

.cp_ball2 {
  animation-delay: .1s;
}

.cp_ball3 {
  animation-delay: .2s;
}

@keyframes ball_typing {
  20% {
    transform: translateY(5px);
  }
  40% {
    transform: translateY(5px);
  }
  80% {
    transform: translateY(-10px);
    opacity: 0.5;
  }
  100% {
    transform: translateY(0);
  }
}
</style>