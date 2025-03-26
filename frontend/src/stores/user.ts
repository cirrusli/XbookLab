import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    isLoggedIn: false,
    user: {
      id: '',
      nickname: '未登录用户',
      avatar: 'https://cdn.example.com/default-avatar.png'
    }
  }),
  actions: {
    login(userData) {
      this.isLoggedIn = true
      this.user = {
        id: userData.id,
        nickname: userData.nickname || '新用户',
        avatar: userData.avatar || this.user.avatar
      }
    },
    logout() {
      this.isLoggedIn = false
      this.user = {
        id: '',
        nickname: '未登录用户',
        avatar: 'https://cdn.example.com/default-avatar.png'
      }
    }
  }
})