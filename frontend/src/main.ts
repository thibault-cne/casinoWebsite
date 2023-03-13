import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import router from "./router/router";
import mitt from 'mitt';

const emitter = mitt();

const app = createApp(App).use(router);
app.config.globalProperties.emitter = emitter;

app.mount("#app");

document.querySelector('html')?.setAttribute('data-theme', "dark");
