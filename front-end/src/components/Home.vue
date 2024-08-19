<script lang="ts" setup>
import axios from "axios";
import { useQuery } from "@tanstack/vue-query";
import { ref, watchEffect, watch } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

type Data = {
  Token: string;
};

const email = ref("");
const password = ref("");

const postLogin: () => Promise<void> | Data = async () => {
  const res = await axios.post("http://localhost:8080/login", {
    email: email.value,
    password: password.value,
  });
  const data = res.data; 
  return data;
};

const showModal = (message: string) => {
  alert(message);
};

const login = async () => {
  if (!email.value || !password.value) {
    showModal("Preencha todos os campos");
    return;
  }
  const data = await postLogin();
  watchEffect(() => {
    if (data) {
      showModal("Login efetuado com sucesso");
      getProfile(data);
    } else showModal("Email ou senha incorretos");
  });
};

const getProfile = async (data: { username: string; token: string }) => {
  const { username, token } = data;

  router.push({
    name: "UserProfile",
    params: { username },
  });
};
</script>

<template>
  <div class="bg-inherit text-white">
    <h1 class="pt-12 ml-4 text-6xl text-orange-600">Curiously</h1>
    <form
      @submit.prevent="login"
      class="flex flex-col justify-center items-center"
    >
      <input
        type="email"
        class="w-1/2 p-2 m-2 text-lg bg-black focus:outline-none"
        placeholder="Email"
        v-model="email"
      />
      <input
        type="password"
        class="w-1/2 p-2 m-2 text-lg bg-black focus:outline-none"
        placeholder="Password"
        v-model="password"
      />
      <p class="text-gray-500">não tem uma conta ?</p>
      <button
        class="w-1/2 p-2 m-2 text-lg bg-orange-600 text-white"
        type="submit"
      >
        Login
      </button>
    </form>
  </div>
</template>
