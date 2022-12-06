<template>
  <h3>Search for posts by Title</h3>
  <my-input
      id="searchBar"
      v-model="searchQuery"
      placeholder="What would you like to search for?"
  />
  <div class="app_btns">
    <my-button @click="showDialog">Create a Post</my-button>
    <my-select v-model="selectedSort" />
  </div>
  <select v-model="filterQuery" class="filter">
    <option disabled value="" >Filter Posts by Category</option>
    <option>🚗 Cars </option>
    <option>🐕 Dogs </option>
    <option>🗺 Travel </option>
    <option value = "">Back to All Posts</option>
  </select>

  <my-dialog v-model:show="dialogVisible"
  ><post-form @create="createPost"
  /></my-dialog>
  <div v-if="posts != null" v-bind:id="{'searchBar.active': isActive}">
    <post-list :posts="searchPosts" @remove="removePost" v-if="!isPostsLoading" />
    <div v-else>LOADING...</div>
  </div>
  <h2 v-else>No posts to display!</h2>
</template>

<script>
import PostForm from "@/components/PostForm";
import PostList from "@/components/PostList";
import MyDialog from "@/components/UI/MyDialog";
import MyButton from "@/components/UI/MyButton";
import axios from "axios";
axios.defaults.withCredentials = true;
import MySelect from "@/components/UI/MySelect";
import MyInput from "@/components/UI/MyInput";
import store from "@/store";
import router from "@/router/router";
import {connection} from "@/components/webSockets";


export default {
  props: ["dontShowDialog"],
  components: {
    MyInput,
    MySelect,
    MyButton,
    MyDialog,
    PostList,
    PostForm,
  },
  data() {
    return {
      posts: [],
      dialogVisible: this.dontShowDialog,
      isPostsLoading: false,
      isActive: false,
      selectedSort: "",
      searchQuery: "",
      filterQuery: "",
      sortOptions: [
        { value: "title", name: "by name" },
        { value: "body", name: "by name" },
      ],
    };
  },

  methods: {
    getFromStore() {
      return store.state.posts
    },
    createPost(post) {
      this.posts.push(post);
      this.dialogVisible = false;
    },
    removePost(post) {
      this.posts = this.posts.filter((p) => p.id !== post.id);
    },
    showDialog() {
      this.dialogVisible = true;
    },
    async fetchPosts() {
      await axios({
            url: "http://localhost:9200/allPosts",
            method: "get",
            headers: {'Content-Type': 'application/json'},
            withCredentials: true,
            origin: true,
            // origin: "http://localhost:8080",
          },
      ).catch((error) => {
        console.log("We got an error: ", error)
      }).then((response) => {
        if (response.data != null && response.data.type === 'Error') {
          store.commit('reset')
          alert("You have been logged out!")
          localStorage.clear()
          router.push('/')
          location.reload();
          connection.close();
        } else {
          store.commit('addPosts', response.data) //add data to store
          this.posts = store.getters.loadedPosts //populate current posts with posts from store
        }
      })
    },
  },

  mounted() {
    this.fetchPosts();
  },
  computed: {
    searchPosts() {
      if (this.searchQuery != "") {
        return this.posts.filter((post) =>
            post.title.toLowerCase().includes(this.searchQuery.toLowerCase())
        )
      }

      if (this.filterQuery != "") {
        return this.posts.filter((post) =>
            post.category.includes(this.filterQuery))
      }
      return this.posts

    },
  }
};
</script>

<style>

#searchBar {
  width: 40%;
}

.filter {
  margin-bottom: 15px;
}

.app_btns {
  margin: 15px 0;
  display: flex;
  justify-content: space-between;
}
.page_wrapper {
  display: flex;
  margin-top: 15px;
}
.page {
  border: 1px solid black;
  padding: 10px;
}
.current-page {
  border: 2px solid teal;
}
.observer {
  height: 30px;
  background: greenyellow;
}
</style>