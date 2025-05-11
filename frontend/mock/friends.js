import Mock from 'mockjs'

export default {
    url: '/api/getFriendsBook',
    method: 'get',
    response: ({ query }) => {
        const { data } = Mock.mock({
            'data|10-20': [{
                'id|+1': 1,
                name: '@cname',
                avatar: Mock.Random.image('100x100', Mock.Random.color(), '#FFF', 'png', '@name'),
                reading: '正在阅读',
                bookTitle: '@ctitle(3, 8)',
                bookCover: () => Mock.Random.image('150x200', '#4A7BF7', '#FFF', 'png', '书'),
                'bookProgress|1-100': 1,
                bookComment: '@cparagraph(1, 3)',
                lastActive: () => Mock.Random.datetime('yyyy-MM-dd HH:mm'),
                type: '@pick(["1","2"])', //1表示最近活跃，2表示同城
            }]
        })

        let filteredList = [...data]
        const type = query.type

        if (type) {
            filteredList = data.filter(
                item => item.type === type
            )
        }

        return {
            code: 200,
            data: data.sort(() => 0.5 - Math.random()), // 随机排序
            total: data.length
        }
    }
}