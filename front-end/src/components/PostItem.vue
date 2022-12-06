<template>
  <div class="post">
    <div>
      <div><strong>Post Title: </strong> {{ post.title }}</div>
      <div><strong>Post Description: </strong> {{ post.body }}</div>
      <div><strong>Post Author: </strong> {{ post.username }}</div>
      <div><strong>Post Date: </strong> {{ post.date }}</div>
      <div><strong>Category: </strong> {{ post.category }}</div>
    </div>
    <div class="post__btns"></div>
    <div v-if="activateDeleteButton()">
      <my-button @click="$emit('remove', post); deletePostFromDB();">Delete</my-button>
    </div>
    <div>
    <my-button @click="storeOnePost()">Read Post</my-button>
  </div>

    </div>
</template>
<script>
import MyButton from "@/components/UI/MyButton";
import store from "@/store";
import axios from "axios";

export default {
  name: "PostItem",
  components: { MyButton },
  props: {
    post: {
      type: Object,
      required: true,
    },
    data() {
      return {
        user: "",
        post: {
          id: "",
          username: "",
          title: "",
          body: "",
          category: "",
          date: "",
        },
        parsedDate: "",
      }
    }
  },
  methods: {
    storeOnePost() {
      store.state.post.id = this.post.id
      store.state.post.username = this.post.username
      store.state.post.title = this.post.title
      store.state.post.body = this.post.body
      store.state.post.category = this.post.category
      store.state.post.date = this.post.date
      this.$router.push({path: `/posts/${this.post.id}`})
    },
    activateDeleteButton() {
    this.user = store.state.user.username
      return this.user === this.post.username
   },
    deletePostFromDB() {
        axios({
          method: "post",
          url: "http://localhost:9200/deletePost",
          data: this.post.id,
          headers: {'Content-Type': 'text/plain'},
          withCredentials: true,
          origin: true,
          //  origin: "http://localhost:8080",
        }).then((result) => {
          this.msg = result.data.msg;
          this.flag = result.data.check;
          if (this.flag === true) {
           // alert(result.data.msg);

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
    }
  },
};
</script>
<style scoped>
.post {
  margin: 15px;
  padding: 15px;
  border: 2px solid teal;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.post__btns {
  white-space: nowrap;
  padding-left: 50px;
  margin-left: auto;
}
</style>
