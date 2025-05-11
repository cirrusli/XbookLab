import { get } from '@/utils/axios';

export function GetGroupListApi(params) {
    return get('/api/getGroups', params);
}