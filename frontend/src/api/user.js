import { get, post } from '@/utils/axios';

// 登录
export function LoginApi(data) {
    return post('/api/auth/login', data);
}

// 注册
export function RegisterApi(data) {
    return post('/api/register', data)
}

// 退出登录
export function LogoutApi(data) {
    return post('/api/auth/logout', data)
}

// 编辑用户信息
export function UpdateUserInfoApi(data) {
    return post('/api/user/update', data);
}

// 修改密码
export function ChangePasswordApi(data) {
    return post('/api/auth/change-password', data);
}

// 忘记密码
export function ForgotPasswordApi(data) {
    return post('/api/forgotPassword', data)
}

// 我的个人信息
export function GetUserProfileApi(data) {
    return post("/api/user/profile", data)
}

// 关注列表
export function GetFollowersListApi() {
    return get('/api/user/following')
}

// 粉丝列表
export function GetFollowMeListApi() {
    return get('/api/user/followers')
}