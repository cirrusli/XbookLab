import { message } from 'ant-design-vue';
import Mock from 'mockjs'

// 随机生成已关注的人
const FollowersList = Mock.mock({
    'data|5-20': [{
        'id|+1': 1,                  // 自增ID
        'name': '@cname',             // 中文姓名
        'avatar': '@image("100x100")', // 随机头像图片URL
    }]
});

// 随机生成我的粉丝
const FollowMeList = Mock.mock({
    'data|5-20': [{
        'id|+1': 1,                  // 自增ID
        'name': '@cname',             // 中文姓名
        'avatar': '@image("100x100")', // 随机头像图片URL
    }]
});

export default [
    {
        url: '/api/login',
        method: 'post',
        response: ({ body }) => {
            const { email, password } = body
            if (email === 'admin@gmail.com' && password === '123456') {
                return {
                    code: 200,
                    data: {
                        token: 'admin-token',
                        name: '管理员',
                        avatar: Mock.Random.image('100x100', Mock.Random.color(), '#FFF', 'png', '@name'),
                        roles: ['admin']
                    },
                    message: '登录成功'
                }
            } else {
                return {
                    code: 401,
                    data: '',
                    message: '用户名或密码错误'
                }
            }
        }
    },

    {
        url: '/api/register',
        method: 'post',
        response: ({ body }) => {
            const { email, password } = body
            if (!email || !password) {
                return {
                    code: 401,
                    data: '',
                    message: '注册失败'
                }
            } else {
                return {
                    code: 200,
                    data: '',
                    message: '注册成功，请登录'
                }
            }

        }
    },

    {
        url: '/api/logout',
        method: 'post',
        response: () => {
            return {
                code: 200,
                data: '',
                message: '退出成功'
            }
        }
    },

    {
        url: '/api/updateUserInfo',
        method: 'post',
        response: ({ body }) => {
            const { name, avatar } = body
            if (name && avatar) {
                return {
                    code: 200,
                    data: '',
                    message: '修改成功'
                }
            } else {
                return {
                    code: 401,
                    data: '',
                    message: '修改失败'
                }
            }
        }
    },
    {
        url: '/api/changePassword',
        method: 'post',
        response: ({ body }) => {
            const { old_password, new_password } = body
            if (old_password === '123456' && new_password) {
                return {
                    code: 200,
                    data: '',
                    message: '修改成功'
                }
            } else {
                return {
                    code: 401,
                    data: '',
                    message: '修改失败'
                }
            }
        }
    },

    {
        url: '/api/forgotPassword',
        method: 'post',
        response: ({ body }) => {
            const { password } = body
            if (password) {
                return {
                    code: 200,
                    data: '',
                    message: '修改成功'
                }
            } else {
                return {
                    code: 401,
                    data: '',
                    message: '修改失败'
                }
            }
        }

    },

    {
        url: '/api/getUserProfile',
        method: 'post',
        response: () => {
            return {
                code: 200,
                data: {
                    id: '',
                    name: '',
                    avatar: '',
                    followCount: FollowersList.data.length,
                    followMeCount: FollowMeList.data.length,
                    topicCount: 10,
                    bookCount: 20
                }
            }
        }
    },

    {
        // 已关注列表
        url: '/api/getFollowersList',
        method: 'get',
        response: () => {
            const { data } = FollowersList
            return {
                code: 200,
                data: {
                    total: data.length,  // 总数量
                    list: data           // 关注者列表
                },
                message: '获取成功'
            };
        }
    },

    {
        // 粉丝列表
        url: '/api/getFollowMeList',
        method: 'get',
        response: () => {
            const { data } = FollowMeList

            return {
                code: 200,
                data: {
                    total: data.length,  // 总数量
                    list: data           // 粉丝列表
                },
                message: '获取成功'
            };
        }
    },

]