import { createStore } from 'vuex'

export interface User {
  id: number
  nickname: string
  avatar: string
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
      state.user.currentUser = payload.user
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