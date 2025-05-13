import { get, post } from '@/utils/axios';

// 获取首页热门话题
export function GetIndexHotTopicApi(params) {
    return get('/api/topic/', params).then(res => {
        return {
            data: res.data.map(item => ({
                id: item.topic_id,
                title: item.title,
                excerpt: item.content,
                likes: item.like_count,
                tag: item.tag_name,
                views: item.views,
                comments: item.comments,
                author: item.author,
                date: item.date,
                timestamp: item.timestamp,
            }))
        };
    });
}

// 获取热门榜单
export function GetHotTopicListApi(params) {
    return get('/api/getHotTopic', params);
}

// 获取话题详情
export function GetTopicDetailApi(id) {
    return get(`/api/topic/${id}`);
}

// 点赞话题
export function LikeTopicApi(params) {
    return post("/api/like/topic", params)
}

// 创建话题
export function CreateTopicApi(params) {
    return post("/api/topic/", params)
}

// 删除话题
export function DeleteTopic(id) {
    return post(`/api/deleteTopic${id}`)
}