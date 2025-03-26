<template>
  <div class="book-card" @mousemove="handleMouseMove">
    <div class="cover-wrapper">
      <img :src="book.cover || '/placeholder.jpg'" alt="Book cover" class="cover" />
    </div>
    <div class="content">
      <h3 class="title">{{ book.title }}</h3>
      <p class="meta">{{ book.author }} · {{ book.publisher }}</p>
      <div class="tags">
        <el-tag 
          v-for="(tag, index) in book.tags.slice(0, 3)" 
          :key="index" 
          size="small"
        >
          {{ tag }}
        </el-tag>
      </div>
      <div class="stats">
        <span class="stat">
          <el-icon><Star /></el-icon>
          {{ book.collects || 0 }}
        </span>
        <span class="stat">
          <el-icon><ChatDotRound /></el-icon>
          {{ book.discussions || 0 }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { Star, ChatDotRound } from '@element-plus/icons-vue';

const props = defineProps({
  book: {
    type: Object,
    required: true
  }
});

const handleMouseMove = (e) => {
  const card = e.currentTarget;
  const xAxis = (window.innerWidth / 2 - e.pageX) / 25;
  const yAxis = (window.innerHeight / 2 - e.pageY) / 25;
  card.style.transform = `rotateY(${xAxis}deg) rotateX(${yAxis}deg)`;
};
</script>

<style scoped>
.book-card {
  width: 100%;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: transform 0.3s ease;
  cursor: pointer;
}

.cover-wrapper {
  aspect-ratio: 4/6;
  overflow: hidden;
}

.cover {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.content {
  padding: 12px;
}

.title {
  font-size: 18px;
  font-weight: bold;
  margin: 0 0 8px 0;
  color: #333;
}

.meta {
  font-size: 12px;
  color: #666;
  margin: 0 0 12px 0;
}

.tags {
  display: flex;
  gap: 6px;
  margin-bottom: 12px;
  flex-wrap: wrap;
}

.stats {
  display: flex;
  gap: 16px;
}

.stat {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 12px;
  color: #666;
}

@media (max-width: 768px) {
  .book-card {
    max-width: 100%;
  }
}
</style>