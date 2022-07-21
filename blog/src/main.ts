import { createApp, h } from "vue";
import App from "@/App.vue";
import router from "./router";
import store from "./store";

createApp({
  render: () => h(App),
})
  .use(router)
  .use(store)
  .mount("#app");
