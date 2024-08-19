<script lang="ts" setup>
import { ref, onMounted, computed} from "vue";
import { useRoute } from "vue-router";
import axios from "axios";

type Question = {
  id: number;
  content: string;
  created_at: string;
  answered: boolean;
  answer: string[] | null;
};

type User = {
  username: string;
  email: string;
  questions: Question[];
  created_at: string;
};

const route = useRoute();
const user = ref<User | null>(null);
const isLoading = ref<boolean>(true);
const error = ref<string | null>(null);

const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  if (isNaN(date.getTime())) {
    return "Invalid date";
  }
  return new Intl.DateTimeFormat("pt-BR", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
    timeZoneName: "short",
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
      }
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
  <div class="bg-neutral-950 min-h-screen p-8">
    <div v-if="user" class="max-w-2xl mx-auto bg-white shadow-lg rounded-lg overflow-hidden">
      <div class="p-6">
        <h1 class="text-3xl font-bold text-gray-800 mb-2">{{ user.username }}</h1>
        <p class="text-sm text-gray-600 mb-6">Membro desde: {{ formatDate(user.created_at) }}</p>
        
        <div v-for="question in user.questions" :key="question.id" class="mb-6 last:mb-0">
          <p class="text-blue-600 font-semibold mb-2">{{ question.content }}</p>
          <div v-if="question.answer && question.answer.length > 0" class="pl-4 border-l-2 border-blue-200">
            <p v-for="answer in question.answer" :key="answer" class="text-red-700 mb-1 last:mb-0">
              {{ answer }}
            </p>
          </div>
          <p v-else class="text-gray-500 italic pl-4"></p>
        </div>
      </div>
    </div>
    <div v-else class="text-center py-8 text-gray-600">
      Loading user data...
    </div>
  </div>
</template>
<!--<template>
  <div class="bg-gray-900 min-h-screen text-white p-8">
    <div v-if="isLoading" class="text-center py-8">
      <div
        class="animate-spin rounded-full h-16 w-16 border-t-2 border-b-2 border-orange-500 mx-auto"
      ></div>
      <p class="mt-4 text-lg">Loading profile...</p>
    </div>
    <div v-else-if="error" class="text-center py-8">
      <p class="text-red-500 text-lg">{{ error }}</p>
    </div>
    <div v-else-if="user" class="max-w-4xl mx-auto">
      <div class="bg-gray-800 rounded-lg shadow-lg p-6 mb-8">
        <h1 class="text-4xl font-bold text-orange-500 mb-2">
          {{ user.username }}
        </h1>
        <p class="text-gray-400">
          Member since: {{ formatDate(user.createdAt) }}
        </p>
      </div>
     <h2 class="text-2xl font-semibold mb-4">Questions & Answers</h2>
      <div v-if="sortedQuestions && sortedQuestions.length > 0">
        <div
          v-for="question in sortedQuestions"
          :key="question.id"
          class="bg-gray-800 rounded-lg shadow-lg p-6 mb-6"
        >
          <div class="mb-4">
            <h3 class="text-xl font-semibold text-orange-400">
              {{ user.value.questions[0].content }}
            </h3>
            <p class="text-sm text-gray-400 mt-1">
              Posted on: {{ formatDate(question.created_at) }}
            </p>
            <p class="text-sm text-gray-400">
              Status: {{ question.answered ? 'Answered' : 'Not answered' }}
            </p>
          </div>
          <div
            v-if="question.answered && question.answer && question.answer.length > 0"
            class="ml-6"
          >
            <h4 class="text-lg font-medium mb-2">Answers:</h4>
            <div
              v-for="(answer, index) in question.answer"
              :key="index"
              class="bg-gray-700 rounded p-4 mb-2"
            >
              <p>{{ answer }}</p>
            </div>
          </div>
          <p v-else-if="question.answered" class="text-gray-500 italic ml-6">No answers provided.</p>
          <p v-else class="text-gray-500 italic ml-6">Not answered yet.</p>
        </div>
      </div>
      <p v-else class="text-gray-500 italic">No questions asked yet.</p>
    </div>
  </div>
  </template> -->