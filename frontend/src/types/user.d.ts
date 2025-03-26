export interface User {
  id: string
  nickname: string
  avatar: string
  email?: string
  bio?: string
}

export interface UserState {
  isLoggedIn: boolean
  user: User
  login: (payload: { 
    id: string; 
    nickname: string; 
    avatar: string; 
    email?: string 
  }) => void
}