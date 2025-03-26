<template>
  <div class="recommend-container">
    <h2 class="section-title">为你推荐</h2>
    <div class="book-list">
      <div v-for="book in books" :key="book.id" class="book-card">
        <img :src="book.cover" class="book-cover" alt="封面" />
        <div class="book-info">
          <h3 class="book-title">{{ book.title }}</h3>
          <div class="book-rating">
            <span class="rating-stars">★★★★☆</span>
            <span class="rating-score">{{ book.rating }}</span>
          </div>
          <p class="book-desc">{{ book.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Book {
  id: number
  title: string
  cover: string
  rating: number
  description: string
}

const books = ref<Book[]>([
  {
    id: 1,
    title: '示例图书1',
    cover: 'https://picsum.photos/200/300',
    rating: 4.5,
    description: '这是一本关于人工智能的入门书籍，适合初学者阅读'
  },
  {
    id: 2,
    title: '示例图书2',
    cover: 'https://picsum.photos/200/300',
    rating: 4.2,
    description: '深入讲解Go语言高级特性的专业书籍'
  }
]) as Ref<Book[]>

import { ref } from 'vue'
import type { Ref } from 'vue'
</script>

<style scoped>
.recommend-container {
  padding: 20px 0;
}

.section-title {
  font-size: 20px;
  color: #333;
  margin-bottom: 30px;
}

.book-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 30px;
}

.book-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  overflow: hidden;
  transition: transform 0.2s;
}

.book-card:hover {
  transform: translateY(-5px);
}

.book-cover {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.book-info {
  padding: 15px;
}

.book-title {
  font-size: 16px;
  margin-bottom: 8px;
}

.book-rating {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.rating-stars {
  color: #ffac2d;
  margin-right: 8px;
}

.rating-score {
  color: #666;
  font-size: 14px;
}

.book-desc {
  color: #666;
  font-size: 14px;
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>