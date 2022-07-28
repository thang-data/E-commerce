import type { AxiosInstance } from "axios"
import type { Router } from "vue-router"
import axiosInstance from "./setting"

export default class Repository {
  router: Router | null = null
  axios: AxiosInstance = axiosInstance

  withRouter(router: Router) {
    this.router = router
    return this
  }
}
