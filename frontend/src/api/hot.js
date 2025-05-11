import { get, post } from '@/utils/axios';

// 获取首页热门话题
export function GetIndexHotTopicApi(params) {
    return get('/api/getIndexHotTopic', params);
}

// 获取热门榜单
export function GetHotTopicListApi(params) {
    return get('/api/getHotTopic', params);
}

// 获取话题详情
export function GetTopicDetailApi(id) {
    return get(`/api/getTopicDetail/${id}`);
}

// 点赞话题
export function LikeTopicApi(params) {
    return post("/api/likeTopic", params)
}

// 创建话题
export function CreateTopicApi(params) {
    return post("/api/createTopic", params)
}

// 删除话题
export function DeleteTopic(id) {
    return post(`/api/deleteTopic${id}`)
}