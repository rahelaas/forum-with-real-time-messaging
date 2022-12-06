<template>
  <div><router-link to="/posts">
    <my-button type="submit" class="btn bt">Back to posts</my-button>
  </router-link></div>
  <br>
  <br>
  <br>
  <br>
  <div>
    <h2>Title: {{ post.title }}</h2>
  </div>

  <div class="post">
    <div><h4>{{post.body}}</h4></div>
    <div class="space"> </div>
    <div class="small">Post by <b>{{post.username}}</b></div>
    <div class="small">Category: <b>{{post.category}}</b></div>
    <div class="small">{{post.date}}</div>
      </div>

  <br>


  <comment-form class="commentForm"></comment-form>
  <br>
  <br>
  <div v-if="comments != null">
    <h3>Comments</h3>
    <div
        v-for="comment in comments"
        :key="comment.id"
        @remove="$emit('remove', comment)">
      <div class="comment">
        <div>{{comment.body}}</div>
       <div class="space"></div>
        <div class="small">Comment by <b>{{comment.username}}</b></div>
         <div class="small">{{comment.date}}</div>
      </div>
    </div>
  </div>
  <h4 v-else>No comments yet, be the first one!</h4>

</template>

<script>

import store from "@/store";
import axios from "axios";
axios.defaults.withCredentials = true;

import CommentForm from "@/components/CommentForm";
import router from "@/router/router";
import {connection} from "@/components/webSockets";
import MyButton from "@/components/UI/MyButton";

export default {
  name: "PostIdPage",
  components: {MyButton, CommentForm, /*MyButton*/},
  data() {
    return{
      comments: [],
      postID: {
        id: "",
      },
      post: {
        title: "",
        body: "",
        username: "",
        category: "",
        date: "",
        id: ""
      },
    }
  },
methods: {

  async fetchPost() {
      let num = Number(this.$route.params.id)
      this.postID.id = num
    let response = await axios({
          url: "http://localhost:9200/singlePost",
          method: "post",
          data: this.postID,
          headers: {'Content-Type': 'application/json'},
          withCredentials: true,
          origin: true,
        },
    ).catch(() => {
      return response
    })
    if (response.data.type === 'Error') {
      store.commit('reset')
      alert("You have been logged out!")
      localStorage.clear()
      await router.push('/')
      location.reload();
      connection.close();

    } else {
      this.post = response.data
    }
  },
  async fetchComments() {
    let num = Number(this.$route.params.id)
    this.postID.id = num

    await axios({
          url: "http://localhost:9200/allComments",
          method: "post",
          data: this.postID,
          headers: {'Content-Type': 'application/json'},
          withCredentials: true,
          origin: true,
          // origin: "http://localhost:8080",
        },
    ).then((response) => {
          let result = response.data
          store.commit('addComments', result) //add data to store
          this.comments = store.getters.loadedComments //populate current posts with posts from store
        }
    ).catch((error) => {
      console.log("error from getting comments ", error)
    })
  },
},
  mounted() {
   this.fetchPost()
    this.fetchComments();

  },
};
</script>

<style scoped>
.post {
  margin-top: 15px;
  margin-bottom: 15px;
  padding: 15px;
  border: 2px solid teal;
  display: inline-list-item;
  align-items: center;
  justify-content: space-between;
}
.comment {
  width: 70%;
  font-size: large;
  margin-top: 15px;
  margin-bottom: 15px;
  padding: 15px;
  border: 1px solid teal;
  display: inline-list-item;
  justify-content: space-between;
}
.commentForm {
  margin-left: 15%;
}
.small {
  font-size: 12px;
  color: #515145;
}
.space {
  margin-top:5px;
}
b {
  color: teal;
}
.bt{
  float: left;
  margin-left: 0px;
  margin-bottom: 10px;
}
</style>
