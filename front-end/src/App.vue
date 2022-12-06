<template>
  <navbar></navbar>

  <div id="content" class="content" >
    <router-view></router-view>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar";
import store from "./store";
import {onUnmounted, provide} from "vue";

export default {
  store,
  components: { Navbar },

  setup() {
    onUnmounted(() => {
      console.log("empty")
      localStorage.removeItem("vuex")
      localStorage.clear()
    }),
    provide('store',store)
  },
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.content {
  padding:20px 50px;
}

label {
  margin-top: 15px;
  margin-bottom: 4px;
  font-family: KindelSerif, serif ;
}
</style>
