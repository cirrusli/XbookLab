import { get, post } from "@/utils/axios"

export function GetManageTagListApi() {
    return get("/api/tag/")
}

export function CreateTagApi(params) {
    return post("/api/tag/", params)
}

export function EditTagApi(params) {
    return post("/api/manage/editTag", params)
}

export function DelTagApi(params) {
    return post('/api/tag/delete', params)
}