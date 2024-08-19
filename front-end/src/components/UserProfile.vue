<script lang="ts" setup>
import { ref } from "vue";
import { useRoute } from "vue-router";

type User = {
  username: string;
  createdAt: string;
  Questions: [];
}

type Questions = [
  {
    id: number;
    content: string;
    createdAt: string;
    answers: string[];
  }
]

const route = useRoute();

const fetchUser = async (username : string) => {
  try {
    const response = await axios.get(`http://localhost:8080/user/${username}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  } catch (error) {
    console.error(error); 
  }
}

watch(
  () => route.params.username,
  async (newUsername) => {
    if (newUsername) {
      await fetchUser(newUsername);
    }
  }, 
  { immediate: true }
);

</script>

<template>
  <div class="flex flex-col bg-inherit">
    <img id="header" src="" alt="header" />
    <div class=" flex flex-col justify-center">
      <img id="profile" src="" alt="profile" />
      <h1 class="text-4xl text-white mt-4">{{ user.username }}</h1>
    </div>
    <div class="flex flex-col justify-center items-center">
      <p class="text-white">Criado em: {{ user.createdAt }}</p>
      <h2 class="text-2xl text-white">Perguntas</h2>
      <div class="flex flex-col justify-center items-center">
        <div v-for="question in user.Questions" :key="question.id">
          <p class="text-white">{{ question.content }}</p>
          <div v-for="answer in question.answers" :key="answer">
            <p class="text-white ml-4">{{ answer }}</p>
          </div>
        </div>
      </div>
      </div>
      </div>
</template>
