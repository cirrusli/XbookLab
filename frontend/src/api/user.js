import { get, post } from '@/utils/axios';

// 登录
export function LoginApi(data) {
    return post('/api/login', data);
}

// 注册
export function RegisterApi(data) {
    return post('/api/register', data)
}

// 退出登录
export function LogoutApi(data) {
    return post('/api/logout', data)
}

// 编辑用户信息
export function UpdateUserInfoApi(data) {
    return post('/api/updateUserInfo', data);
}

// 修改密码
export function ChangePasswordApi(data) {
    return post('/api/changePassword', data);
}

// 忘记密码
export function ForgotPasswordApi(data) {
    return post('/api/forgotPassword', data)
}

// 我的个人信息
export function GetUserProfileApi(data) {
    return post("/api/getUserProfile", data)
}

// 关注列表
export function GetFollowersListApi() {
    return get('/api/getFollowersList')
}

// 粉丝列表
export function GetFollowMeListApi() {
    return get('/api/getFollowMeList')
}