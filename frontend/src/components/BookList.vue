<template>
  <div class="book-list">
    <BookCard 
      v-for="book in books" 
      :key="book.id"
      :book-data="book"
      @click="handleCardClick(book.id)"
    />
  </div>
</template>

<script setup lang="ts">
interface Book {
  id: number
  title: string
  cover: string
  author: string
  press: string
  tags: string[]
  collects: number
  discussions: number
}

import type { PropType } from 'vue'

defineProps({
  books: {
    type: Array as PropType<Book[]>,
    required: true
  }
})

const handleCardClick = (bookId: number) => {
  // 处理图书卡片点击事件
  console.log('Navigate to book detail:', bookId)
}
</script>

<style scoped>
.book-list {
  padding: 24px;
  display: grid;
  gap: 24px;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
}

@media (max-width: 768px) {
  .book-list {
    grid-template-columns: 1fr;
  }
}

.book-card:hover {
  transform: translateY(-3px);
}

.book-cover {
  width: 100%;
  height: 180px;
  object-fit: cover;
  border-radius: 4px;
  margin-bottom: 12px;
}

.book-title {
  font-size: 16px;
  margin-bottom: 8px;
}

.book-meta {
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
  font-size: 14px;
  color: #666;
}

.book-desc {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>