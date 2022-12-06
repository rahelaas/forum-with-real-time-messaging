<template>
  <div v-if="comments.length > 0">
    <h2>List of Posts</h2>
    <transition-group name="user-list"
    ><post-item
        v-for="comment in comments"
        :key="comment.id"
        :post="post"
        @remove="$emit('remove', post)"
    /></transition-group>
  </div>
  <h2 v-else>No comments yet! Be the first one</h2>
</template>

<script>
export default {
  name: "CommentList"
}
</script>

<style scoped>

</style>