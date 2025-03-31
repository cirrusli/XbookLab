import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/views/HomeView.vue'),
    children: [
      {
        path: 'groups',
        name: 'groups',
        component: () => import('@/views/GroupsView.vue')
      },
      {
        path: 'hot',
        name: 'hot',
        component: () => import('@/views/HotView.vue')
      },
      {
        path: 'friends',
        name: 'friends',
        component: () => import('@/views/FriendsView.vue')
      },
      {
        path: 'random',
        name: 'random',
        component: () => import('@/views/RandomView.vue')
      },
      {
        path: '',
        redirect: '/groups'
      },
      {
        path: '/book/:id',
        name: 'book-detail',
        component: () => import('@/views/BookDetail.vue')
      },
      {
        path: '/profile',
        name: 'profile',
        component: () => import('@/views/UserProfile.vue')
      }
    ]
  },
]

export default createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})