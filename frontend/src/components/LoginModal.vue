<template>
  <div class="login-modal">
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label>用户名</label>
        <input v-model="form.username" type="text" required />
      </div>
      <div class="form-group">
        <label>密码</label>
        <input v-model="form.password" type="password" required />
      </div>
      <div class="form-group">
        <label>昵称</label>
        <input v-model="form.nickname" type="text" required />
      </div>
      <div class="form-group">
        <label>邮箱</label>
        <input v-model="form.email" type="email" required />
      </div>
      <div class="form-group">
        <label>个人简介</label>
        <textarea v-model="form.bio" required></textarea>
      </div>
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>
      <div class="button-group">
        <button type="submit" class="submit-btn">登录</button>
        <button type="button" class="register-btn" @click="handleRegister">注册</button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useUserStore } from '@/stores/user'
import type { UserState } from '@/types/user'

const userStore = useUserStore() as UserState
const form = ref({
  username: '',
  password: '',
  nickname: '',
  email: '',
  avatar: 'https://randomuser.me/api/portraits/lego/1.jpg',
  bio: ''
})

const errorMessage = ref('')

import { defineEmits } from 'vue'
const emit = defineEmits(['login-success'])

const handleSubmit = async () => {
  try {
    const { data } = await axios.post('/api/auth/login', {
      username: form.value.username,
      password: form.value.password
    })
    
    localStorage.setItem('token', data.token)
    userStore.login({
      id: data.user.id,
      nickname: data.user.nickname,
      avatar: data.user.avatar,
      email: data.user.email
    })
    emit('login-success', data.user)
    
  } catch (error) {
    errorMessage.value = '用户名或密码错误'
    console.error('登录失败:', error)
  }
}

const handleRegister = async () => {
  try {
    await axios.post('/api/auth/register', {
      username: form.value.username,
      password: form.value.password,
      nickname: form.value.nickname,
      avatar: form.value.avatar,
      bio: form.value.bio,
      email: form.value.email
    })
    errorMessage.value = '注册成功，请登录'
    await handleSubmit()
  } catch (error) {
    errorMessage.value = '注册失败，用户名可能已存在'
    console.error('注册失败:', error)
  }
}

</script>

<style scoped>
.login-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;

  form {
    background: white;
    padding: 2rem;
    border-radius: 8px;
    width: 400px;
    max-width: 90%;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    animation: fadeIn 0.3s ease;
  }

  .form-group {
    margin-bottom: 1.5rem;

    label {
      display: block;
      margin-bottom: 0.5rem;
      font-weight: 500;
    }

    input {
      width: 100%;
      padding: 0.75rem;
      border: 1px solid #ddd;
      border-radius: 4px;
      font-size: 1rem;

      &:focus {
        outline: none;
        border-color: #1890ff;
      }
    }
  }

  .submit-btn {
    width: 100%;
    padding: 0.75rem;
    background: #1890ff;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    cursor: pointer;
    transition: background-color 0.2s;

    &:hover {
      background: #40a9ff;
    }
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(-20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
}
</style>