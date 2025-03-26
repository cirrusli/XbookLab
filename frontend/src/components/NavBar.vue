<template>
  <!-- 配置Element Plus中文语言包 -->
  <el-config-provider :locale="zhCn">
    <div class="navbar">
      <div class="nav-links">
        <router-link v-for="link in navLinks" :key="link.path" :to="link.path">
          {{ link.name }}
        </router-link>
      </div>
      <div class="user-section">
        <el-button v-if="!isLoggedIn" type="primary" @click="handleLogin">
          登录
        </el-button>
        <div v-else class="user-profile">
          <el-avatar :size="40" :src="user.avatar" />
          <span class="username">{{ user.name }}</span>
        </div>
      </div>
    </div>
    <LoginModal 
      v-model="loginVisible"
      @login-success="handleLoginSuccess"
    />
  </el-config-provider>
</template>

<script setup>
import { ref } from 'vue';
import { useUserStore } from '@/stores/user';
import { ArrowRight, User, Fire, Share, Connection } from '@element-plus/icons-vue';
import { storeToRefs } from 'pinia';

const userStore = useUserStore();
const { isLoggedIn, user } = storeToRefs(userStore);

const navLinks = ref([
  { name: '推荐', path: '/' },
  { name: '兴趣小组', path: '/groups' },
  { name: '热榜', path: '/hot' },
  { name: '好友在读', path: '/friends' },
  { name: '随机匹配', path: '/random' },
]);

import LoginModal from './LoginModal.vue';

const loginVisible = ref(false);

const handleLogin = () => {
  loginVisible.value = true;
};

const handleLoginSuccess = (userData) => {
  userStore.login(userData);
  loginVisible.value = false;
};
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 24px;
  background: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.nav-links {
  display: flex;
  gap: 24px;
  align-items: center;
}

.nav-item {
  position: relative;
  padding: 8px 0;
}

.nav-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  width: 24px;
  height: 3px;
  background: #1890ff;
  transform: translateX(-50%);
}

.nav-links a {
  color: #333;
  text-decoration: none;
  font-weight: 500;
}

.nav-links a.router-link-active {
  color: #1890ff;
}

.user-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 8px;
}

.username {
  font-size: 14px;
}
</style>