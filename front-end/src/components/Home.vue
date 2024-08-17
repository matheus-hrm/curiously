<script lang="ts" setup>
import axios from "axios";
import { ref } from "vue";

const email = ref("");
const password = ref("");

const showModal = (message: string) => {
  alert(message);
};

const login = async () => {
  if (!email.value || !password.value) {
    showModal("Preencha todos os campos");
    return;
  }
  await postLogin();
};

const postLogin: () => Promise<void> | undefined = async () => {
  const res = await axios.post("http://localhost:8080/login", {
    email: email.value,
    password: password.value,
  });
  if (res.data) {
    localStorage.setItem("token", res.data.token);
    showModal("Login efetuado com sucesso");
  } else showModal("Email ou senha incorretos");
};
</script>

<template>
  <div class="bg-inherit text-white">
    <h1 class="pt-12 ml-4 text-6xl text-orange-600">Curiously</h1>
    <form @submit.prevent="login" class="flex flex-col justify-center items-center">
      <input type="email" class="w-1/2 p-2 m-2 text-lg bg-black focus:outline-none" placeholder="Email"
        v-model="email" />
      <input type="password" class="w-1/2 p-2 m-2 text-lg bg-black focus:outline-none" placeholder="Password"
        v-model="password" />
      <p class="text-gray-500">não tem uma conta ?</p>
      <button class="w-1/2 p-2 m-2 text-lg bg-orange-600 text-white" type="submit">
        Login
      </button>
    </form>
  </div>
</template>
