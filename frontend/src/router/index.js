import { createRouter, createWebHistory } from "vue-router";
import defaultLayout from "@/layouts/default.vue";
import userAccount from "@/layouts/userAccount.vue";
import userLayout from "@/layouts/user.vue"
const routes = [
    {
        path: '/',
        name: 'Index',
        redirect: '/index',
        component: defaultLayout,
        children: [
            {
                path: '/index',
                name: 'Home',
                component: () => import('@/views/index.vue'),
            },
            {
                path: '/groups',
                name: 'Groups',
                component: () => import('@/views/groups/index.vue'),
            },
            {
                path: '/hot',
                name: 'Hot',
                component: () => import('@/views/hot/index.vue'),
            },
            {
                path: '/friends',
                name: 'Friends',
                component: () => import('@/views/friends/index.vue'),
            }
        ]
    },
    {
        path: '/book',
        name: "Book",
        component: () => import('@/views/book/index.vue'),

    },
    {
        path: '/book/:id',  // 动态参数
        name: 'BookDetail',
        component: () => import('@/views/book/detail.vue'),
        // props: true  // 将路由参数作为props传递
    },
    {
        path: '/topic/:id',
        name: 'TopicDetail',
        component: () => import('@/views/topic/detail.vue'),
    },
    {
        path: '/user',
        name: 'User',
        component: userLayout,
        redirect: '/user',
        children: [
            {
                path: '/user',
                Name: 'User',
                component: () => import('@/views/user/index.vue'),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/user/follow',
                Name: 'UserFollow',
                component: () => import('@/views/user/follow.vue'),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/user/followMe',
                Name: 'UserFollowMe',
                component: () => import('@/views/user/followMe.vue'),
                meta: {
                    needAuth: true
                }
            }
        ],
        meta: {
            needAuth: true
        }
    },
    {
        path: '/user/account',
        name: '',
        component: userAccount,
        redirect: '/user/account',
        children: [
            {
                path: '/user/changePwd',
                name: 'ChangePassword',
                component: () => import('@/views/user/changePwd.vue'),
                meta: {
                    needAuth: true
                },
            },
            {
                path: '/user/account',
                name: 'Account',
                component: () => import('@/views/user/account.vue'),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/manage/book',
                name: 'ManageBook',
                component: () => import('@/views/manage/book/index.vue'),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/manage/book/create',
                name: 'CreateBook',
                component: () => import('@/views/manage/book/create.vue'),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/manage/book/edit/:id',
                name: 'EditBook',
                component: () => import("@/views/manage/book/edit.vue"),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/manage/tag',
                name: 'ManageTag',
                component: () => import('@/views/manage/tag/index.vue'),
                meta: {
                    needAuth: true
                }
            },
            {
                path: '/manage/tag/create',
                name: 'CreateTag',
                component: () => import('@/views/manage/tag/create.vue'),
                meta: {
                    needAuth: true
                }
            },

        ]
    },

    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/user/login.vue'),
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/user/register.vue')
    },
    {
        path: '/forgotPwd',
        name: 'ForgotPwd',
        component: () => import('@/views/user/forgot.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
    scrollBehavior: () => ({ left: 0, top: 0 }),
})

router.beforeEach(async (to, from, next) => {
    const token = localStorage.getItem('token');
    if (!token) {
        if (!to.meta.needAuth) {
            next()
        } else {
            next('/login')
        }
    } else {
        next()
    }
})

export default router;