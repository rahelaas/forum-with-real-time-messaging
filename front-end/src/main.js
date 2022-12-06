import { createApp } from "vue";
import App from "./App";
import components from "@/components/UI";
import router from "@/router/router";
import BootstrapVue3 from 'bootstrap-vue-3'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css'


const app = createApp(App);

components.forEach((component) => {
  app.component(component.name, component);
});
app.use(router)
app.use(BootstrapVue3)
app.mount("#app");


