<template>
  <h1>Welcome to the Forum!</h1>
  <div class="parent">
    <img class="logo"
         :src="require('../assets/logo.png')"
    />
  <form
      @submit.prevent
      @click.prevent
  >

      <br/>
      <br/>
      <h3>Please Log In</h3>
      <p v-if="errors.length">
        <b>Please correct the following error(s):</b>
        <ul>
         <li :key="error" class="error" v-for="error in errors">{{ error }}</li>
        </ul>
      </p>

      <label for="credential">Username / Email</label>
      <my-input
          id="credential"
          v-model="login.credential"
          type="text"
          placeholder="Enter your username or email"
          required
      />

      <label for="password">Password</label>
      <my-input
          id="password"
          v-model="login.password"
          type="password"
          placeholder="Enter your password"
          required
      />

      <br/>

      <div class="buttons">

      <my-button class="btn" type="button" value="Cancel" @click.prevent="reloadPage">Cancel</my-button>
        <my-button class="btn" type="button" value="Send" @click.prevent="SendLogInUser" >Log In</my-button>


      </div>
      <div>
      <strong>Not a Registered Member? </strong>
      <router-link to="/signup">Register</router-link> here
    </div>
    </form>
    </div>
    <br /><br />
  </template>

  <script>
  import MyInput from "@/components/UI/MyInput";
  import MyButton from "@/components/UI/MyButton";
  import axios from "axios";
  //axios.defaults.withCredentials = true;
  import store from "@/store";
  import { connectToWebSocket} from "@/components/webSockets";

  export default {
    components: { MyInput, MyButton },
    name: "LogIn",
    data() {
      return {
        login: {
          credential: "",
          password: "",
          isLoggedIn: false,
        },
        errors: [],
        socket: null,
        info: null
      };
    },
    computed: {
      test() {
        if (store.state.user.isLoggedIn === false) {
          return false
        } else {
          return true
        }
      },
    },
    methods: {
      checkForm: function () {

        this.errors = [];
        if (this.login.credential==="") {
          this.errors.push('Username or email required.');
        }
        if (this.login.password==="") {
          this.errors.push('Password required.');
        }
      },
      reloadPage() {
        window.location.reload();
      },
      SendLogInUser() {
        this.checkForm(this.login)
        if (this.errors.length < 1) {
          axios({
            method: "post",
            url: "http://localhost:9200/",
            data: this.login,
            headers: {'Content-Type': 'text/plain'},
            withCredentials: true,
            origin: true,
          },
          ).then((result) => {
            this.msg = result.data["msg"];
            this.flag = result.data["check"];

            if (this.flag === true) {
              store.state.user.isLoggedIn = result.data["check"]
              store.state.user.username = result.data["username"]
                          connectToWebSocket()

              setTimeout(() => this.$router.push({path: "/posts"}), 200);
            } else {

              alert(result.data["msg"]);
              location.reload()
            }
          });
        }
      },
    },
  };
  </script>

  <style scoped>
  form {
    display: flex;
    flex-direction: column;
    width: 500px;
  }
  .error {
    color: darkred;
  }
  h1 {
    display: flex;
    left: 200px;
  }
  label {
    margin-top: 15px;
    margin-bottom: 4px;
    font-family: KindelSerif, serif;
  }
  .logo {
    height: 200px;
    width: 200px;
  }
  .parent {
    display: flex;
    justify-content: center;
    gap: 30%;
    align-items: center;
  }
  </style>
