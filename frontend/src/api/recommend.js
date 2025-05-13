import { get, post } from '@/utils/axios';
import axios from 'axios';

export function GetBooksListApi(params) {
    return get('/api/book/', params);
}
export function GetRecommendedBooksApi(params) {
    return get('/api/recommend', params);
}
export function GetBookDetailApi(id) {
    return get(`/api/book/${id}`);
}

// 登录状态下调用，记录阅读该书籍行为
export function RecordBookViewApi(params) {
    return post('/api/record/view', params)
}

export function GetBookCommentsApi(params) {
    return get('/api/book/comments', params);
}

// 记录用户评论
export function RecordBookCommentApi(params) {
    return post("/api/comment/", params)
}

// 记录用户评分
export function RecordBookRatingApi(params) {
    return post('/api/record/rating', params)
}

// 后台接口
// 书籍列表
export function GetManageBooksListApi(params) {
    return get('/api/manage/getBooks', params);
}

// 创建书籍
export function CreateBookApi(params) {
    return post('/api/book/', params);
}

// 编辑书籍
export function EditBookApi(params) {
    return post('/api/manage/editBook', params);
}

// 上传图片
export function UploadImgApi(formData) {
    return axios({
        method: "POST",
        url: 'http://localhost:8000/api/book/upload',
        data: formData,
        headers: {
            "Content-Type": "multipart/form-data",
            "Accept": "application/json"
        },
        timeout: 50000
    }).then(response => {
        const { code, message } = response.data;
        if (code === 200) {
            return response.data;
        } else {
            CreateErrorMessage(message || "系统出错");
            return Promise.reject(message || "Error");
        }
    }).catch(error => {
        if (error.response.data) {
            const { message } = error.response.data;
            CreateErrorMessage(message || "系统出错");
        }
        return Promise.reject(error.message);
    });
}