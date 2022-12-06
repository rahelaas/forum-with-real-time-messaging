import { createRouter, createWebHistory } from "vue-router";

import PostPage from "@/pages/PostPage";
import About from "@/pages/About";
import PostIdPage from "@/pages/PostIdPage";
import SignUp from "@/pages/SignUpPage";
import LogIn from "@/components/LogIn";
import PageNotFound from '@/pages/PageNotFound.vue'
import MessagesPage from "@/pages/MessagePage";
import store from "@/store";

const routes = [
  {
    path: "/posts",
    component: PostPage,
    name: "PostPage",
  },
  {
    path: "/about",
    component: About,
  },
  {
    path: "/posts/:id",
    component: PostIdPage,
    name: "PostIdPage",
  },
  {
    path: "/signup",
    component: SignUp,
  },
  {
    path: "/",
    component: LogIn,
    name: "Login",
  },
  {
    path: "/messages",
    component: MessagesPage,
    name: "MessagesPage",
  },
  {
    path: '/:catchAll(.*)*',
    name: "PageNotFound",
    component: PageNotFound,
  },
];

const router = createRouter({
  routes,
  history: createWebHistory(process.env.BASE_URL),
});

router.beforeEach((to, from, next) => {

  const isLoggedIn = store.getters["userStatus"]

  if (!isLoggedIn && to.name == "MessagesPage") {
    return next ("/")
  } else if (!isLoggedIn && to.name == "PostPage") {
    return next ("/")
  } else if (!isLoggedIn && to.name == "PostIdPage") {
    return next("/")
  } else if (!isLoggedIn && to.name == "PageNotFound") {
    return next("/")
  } else if (isLoggedIn && to.path == "/") {
    return next("/posts")
  } else if (isLoggedIn && to.name == "PostIdPage" && store.getters["postIdStatus"] != to.params.id) {
    return next({
      name: "PageNotFound"
    })
  } else {
    next()
  }
})

export default router;