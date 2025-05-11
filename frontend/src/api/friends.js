import { get } from '@/utils/axios';

export function GetFriendsBookListApi(params) {
    return get('/api/getFriendsBook', params);
}