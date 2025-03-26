export interface UserState {
  isLoggedIn: boolean
  user: {
    id: string
    nickname: string
    avatar: string
    email?: string
    bio?: string
  }
  login: (payload: { id: string; nickname: string; avatar: string; email?: string }) => void
}

export const useUserStore: () => UserState

declare module 'pinia' {
  interface PiniaCustomProperties {
    login: (payload: { id: string; nickname: string; avatar: string }) => void
  }
}