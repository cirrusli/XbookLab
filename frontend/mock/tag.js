import Mock from 'mockjs';

const categorys = [
    { id: 1, value: 'literature', tag_name: '文学' },
    { id: 2, value: 'technology', tag_name: '科技' },
    { id: 3, value: 'history', tag_name: '历史' },
    { id: 4, value: 'psychology', tag_name: '心理学' },
    { id: 5, value: 'art', tag_name: '艺术' },
    { id: 6, value: 'business', tag_name: '商业' },
    { id: 7, value: 'philosophy', tag_name: '哲学' },
]

export default [
    {
        url: '/api/manage/getTagList',
        method: 'get',
        response: () => {
            let filteredList = [...categorys]

            return {
                code: 200,
                data: filteredList.sort((a, b) => b.id - a.id),
                total: filteredList.length
            }
        }
    },
    {
        url: '/api/manage/createTag',
        method: 'post',
        response: ({ body }) => {
            const newTag = {
                id: categorys.length + 1,
                tag_name: body.tag_name,
            }

            return {
                code: 200,
                data: newTag,
                message: '创建成功'
            }
        }
    },
    {
        url: '/api/manage/editTag',
        method: 'post',
        response: ({ body }) => {
            const newTag = {
                id: body.id,
                tag_name: body.tag_name,
            }

            return {
                code: 200,
                data: newTag,
                message: '编辑成功'
            }
        }
    },
    {
        url: '/api/manage/delTag',
        method: 'post',
        response: ({ body }) => {
            const id = body.id
            let filteredList = [...categorys]

            filteredList = filteredList.filter((item) => item.id !== id);

            return {
                code: 200,
                data: filteredList,
                message: '删除成功'
            }
        }
    },
]

