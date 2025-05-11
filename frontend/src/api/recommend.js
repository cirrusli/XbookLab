import { get, post } from '@/utils/axios';

export function GetBooksListApi(params) {
    return get('/api/getBooks', params);
}

export function GetBookDetailApi(id) {
    return get(`/api/getBookDetail/${id}`);
}

// 登录状态下调用，记录阅读该书籍行为
export function RecordBookViewApi(params) {
    return post('/api/recordBookView', params)
}

export function GetBookCommentsApi(params) {
    return get('/api/book/comments', params);
}

// 记录用户评论
export function RecordBookCommentApi(params) {
    return post("/api/recordBookComment", params)
}

// 记录用户评分
export function RecordBookRatingApi(params) {
    return post('/api/recordBookRating', params)
}

// 后台接口
// 书籍列表
export function GetManageBooksListApi(params) {
    return get('/api/manage/getBooks', params);
}

// 创建书籍
export function CreateBookApi(params) {
    return post('/api/manage/createBook', params);
}

// 编辑书籍
export function EditBookApi(params) {
    return post('/api/manage/editBook', params);
}

// 上传图片
export function UploadImgApi(params) {
    return post('/api/manage/uploadImg', params);
}