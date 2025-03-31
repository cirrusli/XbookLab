import { defineStore } from 'pinia'

export interface UserState {
  isLoggedIn: boolean
  showLogin: boolean
  user: {
    id: string
    nickname: string
    avatar: string
    email: string
  }
  login: (userData: any) => void
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    isLoggedIn: false,
    showLogin: false,
    user: {
      id: '',
      nickname: '',
      avatar: '',
      email: ''
    },
    login(userData) {
      this.isLoggedIn = true
      this.user = {
        id: userData.id,
        nickname: userData.nickname,
        avatar: userData.avatar,
        email: userData.email
      }
    }
  })
})