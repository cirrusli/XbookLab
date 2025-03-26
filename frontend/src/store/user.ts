import { defineStore } from 'pinia'
import type { User } from '@/types/user'

export const useUserStore = defineStore('user', {
  state: () => ({
    isLoggedIn: false,
    user: {
      id: '',
      nickname: '未登录用户',
      avatar: 'https://cdn.example.com/default-avatar.png',
      bio: '',
      email: ''
    } as User
  }),
  actions: {
    login(userData: { id: string; nickname: string; avatar: string; email?: string }) {  // 更新类型定义
      this.isLoggedIn = true
      this.user = {
        id: userData.id,
        nickname: userData.nickname || '新用户',
        avatar: userData.avatar || this.user.avatar,
        email: userData.email || '',
        bio: this.user.bio  // 保留原有bio
      }
    },
    logout() {
      this.isLoggedIn = false
      this.user = {
        id: '',
        nickname: '未登录用户',
        avatar: 'https://cdn.example.com/default-avatar.png',
        bio: '',
        email: ''
      }
    }
  }
})