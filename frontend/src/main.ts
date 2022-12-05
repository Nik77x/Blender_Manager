import { createApp } from "vue";
import App from "./App.vue";
import "./assets/style.css";
import "./assets/fonts.css";
import "./assets/variables.scss";
import { createPinia } from "pinia";

const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.mount("#app");
