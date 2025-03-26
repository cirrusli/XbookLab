<template>
  <div class="book-detail">
    <h1>{{ bookTitle }}</h1>
    <p>书籍详情内容</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps({
  id: {
    type: String,
    required: true
  }
})

const bookTitle = ref('书籍标题')
</script>

<style scoped>
.book-detail {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}
</style>