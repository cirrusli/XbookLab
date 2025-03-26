<script setup lang="ts">
import { useRoute } from 'vue-router'
const route = useRoute()
</script>

<template>
  <div class="home-container">
    <router-view />
  </div>
</template>

<style scoped>
.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}
.main-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;
  padding: 20px 0;
  border-bottom: 1px solid #eee;
}
.nav-logo {
  font-size: 24px;
  font-weight: 600;
  color: #007bff;
}
.nav-links a {
  margin-left: 30px;
  color: #666;
}
.nav-links a.router-link-exact-active {
  color: #007bff;
}
</style>