import VueRouter from "vue-router";
import Home from "@/pages/Home.vue";
import PageNotFound from "@/pages/PageNotFound.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/csdn", component: Home },
  { path: "*", component: PageNotFound },
];

const router = new VueRouter({
  mode: "history",
  routes,
});

export default router;
