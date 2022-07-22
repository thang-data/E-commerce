import router from "../src/router"
import { createGtm } from "@gtm-support/vue-gtm"

export default createGtm({
  id: process.env.VUE_APP_GTM as string,
  debug: !(process.env.NODE_ENV === "production"),
  vueRouter: router,
})
