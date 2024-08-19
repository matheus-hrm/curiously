import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Profile from "../views/Profile.vue";

const routes = [
  { path: "/", name: "Home", component: Home },
  {
    path: "/user/:username",
    component: Profile,
    name: "Profile",
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
