import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/views/HomeView.vue'),
    children: [
      {
        path: '',
        name: 'recommend',
        component: () => import('@/views/RecommendView.vue')
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