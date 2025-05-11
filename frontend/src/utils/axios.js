import axios from "axios";

import { CreateErrorMessage } from "./alert";

// 创建 axios 实例
const service = axios.create({
    baseURL: '/',
    timeout: 50000,
    // headers: { "Content-Type": "application/;charset=utf-8" },
    // headers: { "Content-Type": "multipart/form-data;charset=utf-8" },
});

// 请求拦截器
service.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        const params = config.data ? config.data : config.params;
        if (token) {
            config.headers.Authorization = token;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截器
service.interceptors.response.use(
    (response) => {
        // console.log(response.data)
        const { code, message } = response.data;
        // console.log(response.data)
        if (code === 200) {
            return response.data;
        } else {
            CreateErrorMessage(message || "系统出错");
            return Promise.reject(message || "Error");
        }
    },
    (error) => {
        // console.error(error)
        if (error.response.data) {
            const { message } = error.response.data;
            CreateErrorMessage(message || "系统出错");
        }
        return Promise.reject(error.message);
    }
);

export function post(url, data = {}, params = {}) {
    return service({
        method: "POST",
        url,
        data,
        params,
    });
}

export function get(url, params = {}) {
    return service({
        method: "GET",
        url,
        params,
    });
}

// 导出 axios 实例
export default service;
