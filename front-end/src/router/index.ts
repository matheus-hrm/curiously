import { createRouter, createWebHistory } from "vue-router";
import Home from "../components/Home.vue";
import UserProfile from "../components/UserProfile.vue";

const routes = [
  { path: "/", component: Home },
  {
    path: "/user/:username",
    component: UserProfile,
    name: "UserProfile",
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
