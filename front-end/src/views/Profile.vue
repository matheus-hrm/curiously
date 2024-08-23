<script lang="ts" setup>
import NavBar from "../components/NavBar.vue";
import QuestionArea from "../components/QuestionArea.vue";
import type { User, Question } from "../types";
import { ref, onMounted, computed, provide } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";

const route = useRoute();
const user = ref<User>(null);
const isLoading = ref<boolean>(true);
const error = ref<string | null>(null);

provide("user", user);

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  if (isNaN(date.getTime())) {
    return "Invalid date";
  }
  return new Intl.DateTimeFormat("pt-BR", {
    year: "numeric",
    month: "long",
  }).format(date);
};

const fetchUser = async () => {
  const token = localStorage.getItem("token");
  if (!token) {
    error.value = "Token not found";
    isLoading.value = false;
    return;
  }

  try {
    const response = await axios.get(
      `http://localhost:8080/user/${route.params.username}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      },
    );
    user.value = response.data;
    if (user.value) {
      console.log(user.value.questions);
    }
  } catch (error) {
    error.value = "Error fetching user data";
    console.error(error);
  } finally {
    isLoading.value = false;
  }
};

const sortedQuestions = () => {
  return user.value?.Questions.sort((a, b) => {
    return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime();
  });
};

onMounted(() => {
  fetchUser();
});
</script>
<template>
  <NavBar />
  <div
    v-if="user"
    class="bg-[#242421] w-full min-h-screen p-8 flex flex-col md:flex-row justify-evenly"
  >
    <div class="bg-white rounded-md space-y-4 w-[450px] h-96">
      <div class="flex flex-col w-full">
        <div id="header" class="bg-zinc-500 w-full h-32 rounded-t">
          <!--TODO: Switch this div to img tag-->
        </div>
        <div class="flex flex-col px-6 top-[-54px] relative">
          <img
            src="https://avatars.githubusercontent.com/u/77426593?v=4"
            alt="avatar"
            class="w-24 rounded-full"
          />
          <div class="flex flex-col justify-evenly ml-2">
            <h1 class="text-3xl font-bold text-gray-800 py-6">
              {{ user?.username }}
            </h1>
            <p class="text-sm text-gray-600">
              Membro desde {{ formatDate(user?.created_at) }}
            </p>
          </div>
        </div>
      </div>
    </div>
    <div
      class="ml-8 w-[1080px] bg-white shadow-lg rounded-lg overflow-hidden h-1/2"
    >
      <div class="p-6">
        <QuestionArea />
        <div
          v-for="question in user.questions"
          :key="question.id"
          class="mb-6 last:mb-0"
        >
          <div class="flex flex-row justify-between">
            <p class="text-zinc-600 font-semibold mr-4 mb-2">
              {{ question.content }}
            </p>
            <p class="text-zinc-500 ml-2">
              {{ formatDate(question.created_at) }}
            </p>
          </div>
          <div
            v-if="question.answer && question.answer.length > 0"
            class="pl-4 border-l-2 border-black"
          >
            <p
              v-for="answer in question.answer"
              :key="answer"
              class="text-gray-500 mb-1 last:mb-0"
            >
              {{ answer }}
            </p>
          </div>
          <p v-else class="text-gray-500 italic pl-4"></p>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="text-center py-8 text-gray-600">Loading user data...</div>
</template>
