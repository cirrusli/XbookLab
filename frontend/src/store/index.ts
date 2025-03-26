import { createStore } from 'vuex'

export interface User {
  id: number
  username: string
  nickname: string
  avatar: string
  bio?: string
  email?: string
}

export interface State {
  user: {
    isLoggedIn: boolean
    currentUser: User | null
    token: string
  }
}

export default createStore<State>({
  state: {
    user: {
      isLoggedIn: false,
      currentUser: null,
      token: ''
    }
  },
  mutations: {
    login(state: State, payload: { user: User; token: string }) {
      state.user.isLoggedIn = true
      state.user.currentUser = {
      ...payload.user,
      username: payload.user.username || '',
      email: payload.user.email || '',
      bio: payload.user.bio || ''
    }
      state.user.token = payload.token
    },
    logout(state: { user: { isLoggedIn: boolean; currentUser: null } }) {
      state.user.isLoggedIn = false
      state.user.currentUser = null
    }
  },
  actions: {},
  modules: {}
})