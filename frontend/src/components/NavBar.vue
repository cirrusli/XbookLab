<script setup>
import { ref } from 'vue';
import { useUserStore } from '@/stores/user';
import { ArrowRight, User, Fire, Share, Connection } from '@element-plus/icons-vue';
import { storeToRefs } from 'pinia';

const userStore = useUserStore();
const { isLoggedIn, user } = storeToRefs(userStore);

const navItems = ref([
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

<template>
  <div class="navbar">
    <div class="nav-links">
      <router-link 
        v-for="item in navItems" 
        :key="item.value"
        :to="item.path"
        class="nav-item"
        :class="{ active: activeTab === item.value }"
      >
        {{ item.name }}
      </router-link>
    </div>
    
    <div class="user-section">
      <template v-if="userStore.isLoggedIn">
        <div class="user-profile">
          <el-avatar :size="32" :src="userStore.avatar" />
          <span class="username">{{ userStore.nickname }}</span>
        </div>
      </template>
      <template v-else>
        <ElButton type="primary" size="small" @click="userStore.showLogin = true">
          登录
        </ElButton>
      </template>
    </div>
  </div>
</template>

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
  color: #333;
  text-decoration: none;
  font-weight: 500;
}

.nav-item.active {
  color: #1890ff;
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