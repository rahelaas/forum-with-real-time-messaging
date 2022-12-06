<template>
  <div>
<h4>Leave a Comment</h4>
    <div class="space"></div>
<form @submit.prevent>
<my-input id="comment" v-model.trim="comment.body" type="text" placeholder="your comment" autocomplete="off" required />
<my-button class="btn" value="Send" @click="createComment()">Create a Comment</my-button>
</form>

</div>
</template>

<script>
import store from "@/store";
import axios from "axios";

export default {
  data() {
    return {
      comment: {
        post_id: "",
        body: "",
        username: store.state.user.username,
      }
    }
  },
  methods: {
    getFromStore() {
      return store.state.posts
    },
    createComment() {
      if (this.comment.body !== "" && this.comment.username !== "") {

        let num = Number(this.$route.params.id)
        this.comment.post_id = num
        axios({
          method: "post",
          url: "http://localhost:9200/comment",
          data: this.comment,
          headers: {'Content-Type': 'text/plain'},
          withCredentials: true,
          origin: true,
          // origin: "http://localhost:8080",
        }).then((result) => {
          this.msg = result.data.msg;
          this.flag = result.data.check;
          if (this.flag === true) {
           // alert(result.data.msg);
            setTimeout(() => {
              this.$router.go({
                name: "allPosts",
                //params: { id: this.$route.params.id },
              });
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
      }
    }


  }
};
</script>

<style scoped>
.space {
  margin-top:5px;
}
</style>
