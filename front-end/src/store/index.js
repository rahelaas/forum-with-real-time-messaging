import Vuex from "vuex";
import VuexPersistence from 'vuex-persist';
const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
});
export default new Vuex.Store({

   plugins: [vuexLocal.plugin],
                state: {
                user: {
                        isLoggedIn: false,
                        username: '',
                    firstName: "",
                    isChattingWith: "",
                    ws: null
                },
                    posts: [],
                    tenMessages: [],
                    messagesLoaded: 0,

                    post: {
                    id: "",
                        username: "",
                        title: "",
                        body: "",
                        date: ""
                    },

                    comments: [],
                    componentKey: null,
                    usersOnline: [],
                    isTyping: false,
                    typer: "",
        },
// getters are like computed properties. They cache their results and only re-evaluate when a dependency is modified.
        getters: {
           loadedPosts(state) {
                return state.posts
            },
            loadedComments(state) {
                return state.comments
            },
            userStatus(state) {
               return state.user.isLoggedIn
            },
            postIdStatus(state) {
                return state.post.id
            },
            getOnlineUsers(state) {
               return state.usersOnline
            }
        },
        mutations: {
                changePassword (state, payload) {
                        state.user.password = payload.newName
                },
            resetState(state) {
                state.user.isLoggedIn = false
            },
            addPosts(state,payload){
                state.posts = payload
            },
            add(state,payload){
                state.componentKey = payload
            },
            addComments(state,payload){
                state.comments = payload
            },
            changeUserOnlineList(state, payload) {
                state.usersOnline = payload
            },
            clearIsChattingWith(state) {
                    state.user.isChattingWith = ""
            },
            clearUserOnlineList(state) {
                    state.usersOnline = []
            },
            changeTenMessages(state, payload) {
                    state.tenMessages = payload
            },
            reset(state) {
                state.user.isLoggedIn = false
            },
            changeMessagesLoaded(state, payload) {
                    state.messagesLoaded = payload
            },
            resetIsTyping(state) {
                state.isTyping = false
            },
            clearTyper(state) {
                state.typer = ""
            },
        },
});