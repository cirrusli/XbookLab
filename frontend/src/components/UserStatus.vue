<template>
  <div class="user-status">
    <div v-if="isLoggedIn" class="user-info">
      <img :src="user.avatar" class="user-avatar" />
      <span class="user-name">{{ user.nickname }}</span>
    </div>
    <button v-else class="login-btn" @click="handleLogin">登录</button>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import  useStore  from '../store'

const store = useStore()
const isLoggedIn = computed(() => store.state.user.isLoggedIn)
const user = computed(() => store.state.user.currentUser)

const handleLogin = () => {
  // 后续接入登录逻辑
  store.commit('user/login', {
    id: 1,
    nickname: '测试用户',
    avatar: 'https://randomuser.me/api/portraits/lego/1.jpg'
  })
}
</script>

<style scoped>
.user-status {
  display: flex;
  align-items: center;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  margin-right: 8px;
}

.user-name {
  font-size: 14px;
  color: #333;
}

.login-btn {
  padding: 8px 16px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>