import VueRouter from "vue-router";
import News from "@/pages/News.vue";
import Home from "@/pages/Home.vue";
import PageNotFound from "@/pages/PageNotFound.vue";

const routes = [
  { path: "/news/:newsType", component: News, props: true },
  { path: "/", component: Home },
  { path: "*", component: PageNotFound },
];

const router = new VueRouter({
  mode: "history",
  routes,
});

export default router;
