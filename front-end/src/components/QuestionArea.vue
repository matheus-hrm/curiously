<script setup lang="ts">
import { ref, inject } from "vue";
import type { User } from "../types/types.ts";
import axios from "axios";

const question = ref("");
const user: User = inject("user");

const sendQuestion = (question: string) => {
  if (question) {
    try {
      axios.post("http://localhost:8080/question", {
        //TODO: change from userid to username
        userId: userId,
        content: question,
        isanonymous: true,
      });
    } catch (err) {
      console.error(err);
    }
  }
  question.value = "";
};
</script>
<template>
  <div class="flex flex-col items-center mb-4">
    <textarea
      name="sendQuestion"
      class="w-full h-24 rounded-md border border-transparent p-2 ring-transparent outline-none text-black focus:border-black focus:border-2 transition duration-300 ease-linear resize-none"
      placeholder="Faça uma pergunta... "
      @input="question = $event.target.value"
      v-on:keydown.enter="sendQuestion(question)"
    />
    <button
      class="w-2/3 h-14 bg-teal-700 text-white p-2 rounded-md mt-2"
      @click="sendQuestion(question)"
    >
      <p class="font-extrabold">Enviar</p>
    </button>
  </div>
</template>
