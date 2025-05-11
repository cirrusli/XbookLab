<template>
  <div>
    <Header />
    <main class="main-container">
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup>
import Header from '@/components/Header/index.vue';
</script>

<style lang="scss" scoped>
.main-container {
  display: grid;
  grid-template-columns: 200px 1fr;
  gap: 24px;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  align-items: start;
}

.content {
  width: 100%;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .groups-container {
    grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  }

  .main-container {
    grid-template-columns: 1fr;
  }
}
</style>
