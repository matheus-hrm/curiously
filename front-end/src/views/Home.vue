<script lang="ts" setup>
import axios from "axios";
import { useQuery } from "@tanstack/vue-query";
import { ref, watchEffect, watch } from "vue";
import { useRouter } from "vue-router";
import type { User } from "../types/types.ts";
import NavBar from "../components/NavBar.vue";

const email = ref("");
const senha = ref("");
const user = ref<User | null>(null);
const error = ref<string | null>(null);
const isLoading = ref<boolean>(false);
const router = useRouter();

const showModal = (message: string) => {
  alert(message);
};

const postLogin: () => Promise<void> | string = async () => {
  const res = await axios.post("http://localhost:8080/login", {
    email: email.value,
    senha: password.value,
  });
  const data = res.data;
  return data;
};

const login = async () => {
  if (!email.value || !senha.value) {
    showModal("Preencha todos os campos");
    return;
  }
  const data = await postLogin();
  if (data) {
    localStorage.setItem("token", data.Token);
    getProfile(data);
  } else showModal("Email ou senha incorretos");
};

const getProfile = async (data: { username: string; token: string }) => {
  const { username, token } = data;
  router.push({
    name: "Profile",
    params: { username },
  });
};
</script>

<template>
  <div class="text-white space-y-5 bg-[#242421] h-full">
    <NavBar />
    <form
      @submit.prevent="login"
      class="flex flex-col justify-center items-center"
    >
      <input
        type="email"
        class="w-1/3 p-2 m-2 mb-4 mt-40 rounded-md text-lg border-2 border-transparent focus:outline-none bg-zinc-800 focus:rounded-none focus:border-2 focus:border-b-white transition duration:200 ease-in-out"
        placeholder="Email"
        v-model="email"
      />
      <input
        type="senha"
        class="w-1/3 p-2 m-2 mt-2 text-lg focus:outline-none border-2 border-transparent bg-zinc-800 focus:rounded-none focus:border-2 focus:border-b-white transition duration:200 ease-in-out"
        placeholder="Senha"
        v-model="senha"
      />
      <div>
        <p class="text-gray-200 my-2">
          não tem uma conta ?
          <router-link to="/register">
            <button class="text-white hover:text-teal-700">registre-se</button>
          </router-link>
        </p>
      </div>
      <button
        class="w-1/3 rounded-md p-2 mt-3 text-lg bg-emerald-400 text-white"
        type="submit"
      >
        Login
      </button>
    </form>
  </div>
</template>
