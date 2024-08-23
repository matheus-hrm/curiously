<template>
  <div class="bg-[#242421] p-8">
    <div class="flex flex-row justify-between mb-6">
      <h1 class="text-4xl ml-4">Curiously</h1>
      <div class="flex flex-row space-x-2 justify-center items-center">
        <button
          class="flex flex-row bg-inherit text-white text-opacity-95 border-2 border-transparent text-sm pr-14 hover:border-b-white hover:border-2 hover:border-opacity-50 hover:text-opacity-100 transition trasnform duration-500 ease-in-out"
          @click=""
        >
          <p class="border border-white rounded-md p-1 px-2 m-2">Ctrl</p>
          <p class="py-1 my-2">+</p>
          <p class="border border-white rounded-md px-2 py-1 m-2">k</p>
          <p class="py-1 my-2">to search</p>
        </button>
        <button
          class="text-[#5b5550] flex justify-center items-center m-4 pl-8"
        >
          <p class="text-2xl">*</p>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>
