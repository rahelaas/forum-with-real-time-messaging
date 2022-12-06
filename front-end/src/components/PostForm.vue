<template>
  <form id="form" @submit.prevent>

    <h3>Your New Post</h3>
    <label for="title">Title</label>
    <my-input id="title" v-model.trim="post.title" type="text" placeholder="Title of the post" autocomplete="off" required />

    <label for="post">Post</label>
    <textarea id="post" v-model.trim="post.body"  placeholder="Write your post..." autocomplete="off" required />
    <br>
    <select v-model="post.category">
      <option disabled value="" >Please Select Category</option>
      <option>🚗 Cars</option>
      <option>🐕 Dogs</option>
      <option>🗺 Travel</option>
    </select>
    <my-button class="btn" value="Send" @click="createPost" >Submit</my-button>
  </form>
</template>

<script>
import MyButton from "@/components/UI/MyButton";
import MyInput from "@/components/UI/MyInput";
import axios from "axios";
import store from "@/store";

export default {
  name: "PostForm",
  components: { MyInput, MyButton},
  data() {
    return {
      dontShowDialog: true,
      post: {
        date: Date.now(),
        username: store.state.user.username,
        title: "",
        body: "",
        category: "",
        flag: true
      },
    };
  },
  methods: {
    createPost() {
      if (this.post.body !== "" && this.post.title !== "" && this.post.category !== "") {
      axios({
        method: "post",
        url: "http://localhost:9200/post",
        data: this.post,
        headers: {'Content-Type': 'text/plain'},
        withCredentials: true,
        origin: true,
        // origin: "http://localhost:8080",

      }).then((result) => {
        this.msg = result.data.msg;
        this.flag = result.data.check;
        if (this.flag === true) {
        //  alert(result.data.msg);

          setTimeout(() => {
            location.reload();
            this.post = {
              flag: true,
              id: "",
              username: "",
              title: "",
              body: "",
            }
          }, 200);
        } else {
          alert(result.data.msg);
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
      this.dontShowDialog = false
    }
    },
    },
};
</script>


<style scoped>
form {
  width: 500px;
  display: flex;
  flex-direction: column;
}
label {
  margin-top: 10px;
  margin-bottom: 4px;
  font-family: KindelSerif, serif;
}
#post {
  height: 140px;
  border: 1px solid teal;
  padding: 15px 15px;
  font-family: Arial, sans-serif;
}

::placeholder{
  font-family: KindelSerif, serif;
}

</style>
