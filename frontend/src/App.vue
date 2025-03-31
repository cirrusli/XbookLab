<template>
  <div id="app">
    <NavBar />
    <router-view/>
    <LoginModal 
      v-model="loginVisible"
      @login-success="handleLoginSuccess"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import NavBar from '@/components/NavBar.vue';
import LoginModal from '@/components/LoginModal.vue';

const userStore = useUserStore();
const { isLoggedIn } = storeToRefs(userStore);

const loginVisible = ref(false);

// 全局路由守卫检查登录状态
router.beforeEach((to) => {
  if (to.meta.requiresAuth && !isLoggedIn.value) {
    loginVisible.value = true;
    return { path: '/' };
  }
});

const handleLoginSuccess = () => {
  loginVisible.value = false;
};
</script>

<style>
#app {
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  min-height: 100vh;
  background-color: #f5f5f5;
}
</style>