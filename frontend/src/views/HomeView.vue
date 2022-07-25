<script setup lang="ts">
import {onMounted, reactive} from "vue";
import Response from "@/components/Response.vue";

  const data = reactive({
    question: '',
    responses: []
  })

  onMounted(()=>{
    console.log(import.meta.env.VITE_API_URL+"/api/day")
    fetch(import.meta.env.VITE_API_URL+"/api/day", ).then((response) => response.json())
        .then((prompt) => {
          data.question=prompt.text
          data.responses = prompt.responses
        });
  })
</script>

<template>
  <main class="p-10">
    <h1 class="text-4xl my-5 font-serif font-bold">A ghoulish AI is asked:</h1>
    <h1 class="text-4xl font-serif">{{ data.question.replace('reddit', 'everyone') }}</h1>
    <Response v-for="response in data.responses" :response="response"></Response>
  </main>
</template>
