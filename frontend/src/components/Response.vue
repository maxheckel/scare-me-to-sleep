<template>
  <div class="my-10 text-xl" @click="visible.show = !visible.show">
    <div class="font-bold">
      AI Pseudonym: {{name}}
    </div>
    <div class="leading-10" v-if="visible.show" v-html="response.text.replace('\n', '<br>')"></div>
    <pre class="text-sm" v-if="!visible.show">Hidden</pre>
  </div>
</template>

<script setup lang="ts">
import { faker } from '@faker-js/faker';
import {reactive} from "vue";
const name = faker.name.findName();
let visible = reactive({
  show: true
});
const props = defineProps(['response'])
</script>

<style scoped>

</style>