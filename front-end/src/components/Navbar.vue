<template>
  <div class="navbar" v-if="userLoggedIn === false">
    <div @click="$router.push('/')"></div>
    <div class="navbar_btns">
      <my-button style="margin-left: 20px" @click="$router.push('/about')"
    >About</my-button>
      <my-button @click="$router.push('/signup')">SignUp</my-button>
      <my-button @click="$router.push('/')">Login</my-button>
</div>
  </div>
    <div class="navbar" v-else>
      <div @click="userLoggedIn">Welcome to Real-Time-Forum:
        <div class="usergreeting">{{userGreeting}}!</div></div>
      <div class="navbar_btns">
        <my-button @click="logOut" >LogOut</my-button>

        <my-button @click="$router.push('/messages')">My Messages</my-button>
        <my-button @click="$router.push('/posts')">Forum Posts</my-button>
      </div>
  </div>
</template>

<script>

import store from "@/store";
import router from "@/router/router";
import axios from "axios";
import {connection, connectToWebSocket} from "@/components/webSockets";

export default {
  // eslint-disable-next-line
  name: "Navbar",
  data() {
    return {
      user: {
        usernameOffline: ""
      },
    }
  },
  computed: {
    userLoggedIn() {
      if (store.state.user.isLoggedIn === false) {
        return false
      } else {
        return true
      }
    },
    userGreeting() {
      return store.state.user.username
    }
  },
  methods: {
     logOut() {
       if (connection !== null) {
         let messageToBackend = {type: "userLeft", sender: store.state.user.username}
         // 0 opening, 1 open, 2 closing, 3 closed
         console.log("Websocket state before sending", connection.readyState)
         connection.send(JSON.stringify(messageToBackend))
       }
      this.loggingOut()
    },

    loggingOut() {

      store.commit('clearUserOnlineList')
      this.user.usernameOffline = store.state.user.username
      axios({
        method: "post",
        url: "http://localhost:9200/logout",
        data: this.user,
        headers: {'Content-Type': 'text/plain'},
        withCredentials: true,
        origin: true,
        //  origin: "http://localhost:8080",
      }).then(() => {
        setTimeout(() => {
          store.commit('reset'),
              store.state.user.isLoggedIn = false
              localStorage.clear()
              router.push('/')
              location.reload();
            }, 200);
      })
    },

    checkWebsocketConnection() {
if ( store.state.user.isLoggedIn) {
  if (connection === null ) {
    console.log("new connection needed")
    connectToWebSocket()
  }
  // 0 opening, 1 open, 2 closing, 3 closed
  console.log("Websocket state on loading the page", connection.readyState)
}


    }
  },
  mounted() {
    this.checkWebsocketConnection()
  }
}
</script>

<style scoped>
.navbar {
  height: 77px;
  background-color:hsla(120,100%,25%,0.15);
  box-shadow: 2px 2px 4px gray;
  display: flex;
  align-items: center;
  padding: 0 30px;
  min-width: 950px;
}
.navbar_btns {
  margin-left: auto;
}
.username {
  margin-left: auto;
}
.usergreeting {

  font-style: italic;
  font-family: Avenir;
  font-weight: bold;
  color: indigo;

}
</style>
