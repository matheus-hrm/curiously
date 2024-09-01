<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";
import NavBar from "../components/NavBar.vue";

const username = ref("");
const email = ref("");
const senha = ref("");
const router = useRouter();

const register = async () => {
  if (!username.value || !email.value || !senha.value) {
    alert("Preencha todos os campos");
    return;
  }
  const res = await axios.post("http://localhost:8080/register", {
    email: email.value,
    username: username.value,
    senha: password.value,
  });
  const data = res.data;
  if (data) {
    localStorage.setItem("token", data.Token);
    router.push({
      name: "Profile",
      params: { username: data.username },
    });
  }
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
        type="username"
        class="w-1/3 p-2 m-2 mb-4 mt-40 rounded-md text-lg border-2 border-transparent focus:outline-none bg-zinc-800 focus:rounded-none focus:border-2 focus:border-b-white transition duration:200 ease-in-out"
        placeholder="Nome de usuário"
        v-model="username"
      />
      <input
        type="email"
        class="w-1/3 p-2 m-2 mb-4 rounded-md text-lg border-2 border-transparent focus:outline-none bg-zinc-800 focus:rounded-none focus:border-2 focus:border-b-white transition duration:200 ease-in-out"
        placeholder="Email"
        v-model="email"
      />
      <input
        type="password"
        class="w-1/3 p-2 m-2 mt-2 text-lg focus:outline-none border-2 border-transparent bg-zinc-800 focus:rounded-none focus:border-2 focus:border-b-white transition duration:200 ease-in-out"
        placeholder="Senha"
        v-model="senha"
      />
      <div class="my-2">
        <router-link to="/" class="text-gray-200 hover:text-white">
          Já tem uma conta?
          <button class="text-white hover:text-teal-700">Faça login</button>
        </router-link>
      </div>
      <button
        class="w-1/3 rounded-md p-2 mt-3 text-lg bg-emerald-400 text-white"
        type="submit"
      >
        Criar conta
      </button>
    </form>
  </div>
</template>
