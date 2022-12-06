<template>
  <div >
    <h2>List of Posts</h2>
    <transition-group name="user-list" :key="componentKey"
    ><post-item
        v-for="post in posts"
        :key="post.id"
        :post="post"
        @remove="$emit('remove', post)"
    /></transition-group>
  </div>
</template>

<script>
import PostItem from "@/components/PostItem";
import store from "@/store";
export default {
  name: "PostList",
  components: { PostItem },
  props: {
    posts: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      componentKey: store.state.componentKey
    }
  }
};
</script>

<style scoped>
.user-list-item {
  display: inline-block;
  margin-right: 10px;
}
.user-list-enter-active,
.user-list-leave-active {
  transition: all 0.6s ease;
}
.user-list-enter-from,
.user-list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

@keyframes s3 {to{transform: rotate(1turn)}}
</style>
