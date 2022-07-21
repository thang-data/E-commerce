import { nextTick } from "vue";
import { createRouter, RouteRecordRaw } from "vue-router";


export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/pages/Home.vue"),
    meta: { title: "Home" },
  },
  {
    {
      path: "/",
      name: "About",
      component: () => import("../views/pages/About.vue"),
      meta: { title: "About" },
    }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.afterEach((to, from) => {
  nextTick(() => {
    document.title = to.meta.title + " | Blogs" || ("Blogs" as any);
  });
});

export default router