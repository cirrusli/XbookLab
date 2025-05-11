import Mock from 'mockjs'
import dayjs from 'dayjs'

const { data } = Mock.mock({
    'data|30-50': [{  // 生成30-50条随机数据
        'id|+1': 1,
        'rank|+1': 1,
        title: '@ctitle(10, 20)',
        excerpt: '@ctitle(40, 60)',
        author: '@cname',
        'views|1000-5000': 1,
        'likes|50-300': 1,
        isLiked: function () {
            // 点赞数高的文章更可能被用户点赞
            return this.likes > 200 ? (Mock.mock('@integer(1, 10)') <= 7 ? 1 : 2)
                : (Mock.mock('@integer(1, 10)') <= 3 ? 1 : 2);
        },
        'comments|10-100': 1,
        // category: '@pick(["literature","technology","history","psychology","art","business","philosophy"])',
        tag: '@pick(["文学","科技","历史","心理学","艺术","商业","哲学"])',
        date: () => dayjs()
            .subtract(Math.floor(Math.random() * 90), 'day') // 生成最近90天的数据
            .format('YYYY-MM-DD'),
        timestamp: () => dayjs()
            .subtract(Math.floor(Math.random() * 90), 'day')
            .valueOf() // 同时存储时间戳便于筛选
    }]
})

export default [
    {
        url: "/api/getIndexHotTopic",   // 首页热门话题
        method: "get",
        response: ({ query }) => {
            let filteredList = [...data]
            const tag = query.tag || ''


            if (tag) {
                filteredList = data.filter(
                    item => item.tag === tag
                )
            }

            return {
                code: 200,
                data: filteredList.sort((a, b) => b.likes - a.likes).slice(0, 6), // 按likes取前6条
                message: '获取成功'
            }
        }
    },
    {
        url: '/api/getHotTopic',
        method: 'get',
        response: ({ query }) => {
            // 根据筛选参数过滤数据
            let filteredData = [...data]
            const range = query.range || ''
            if (range) {
                const now = dayjs()
                filteredData = data.filter(item => {
                    const itemDate = dayjs(item.date)
                    switch (range) {
                        case 'today':
                            return itemDate.isSame(now, 'day')
                        case 'week':
                            return itemDate.isAfter(now.subtract(1, 'week'))      // 最近一周
                        case 'month':
                            return itemDate.isAfter(now.subtract(1, 'month'))   //最近一个月
                        default:
                            return true
                    }
                })
            }

            return {
                code: 200,
                data: filteredData.sort((a, b) => b.timestamp - a.timestamp), // 按时间倒序
                total: filteredData.length
            }
        }
    },

    {
        url: '/api/getTopicDetail/:id',
        method: 'get',
        response: ({ query }) => {
            const id = parseInt(query.id);
            let filteredData = [...data]
            console.log(filteredData)
            const foundItem = data.find(item => item.id === id);

            if (foundItem) {
                return {
                    code: 200,
                    data: foundItem,
                    message: '获取成功'
                }
            } else {
                return {
                    code: 401,
                    data: '',
                    message: '未查询到该id数据'
                }
            }

        }
    },
    {
        url: '/api/likeTopic',
        method: 'post',
        response: ({ body }) => {
            const likeFlag = body.is_liked === 1 ? true : false
            if (likeFlag) {
                return {
                    code: 200,
                    data: '',
                    message: '点赞成功'
                }
            } else {
                return {
                    code: 200,
                    data: '',
                    message: '取消点赞'
                }
            }

        }
    },
    {
        url: '/api/createTopic',
        method: 'post',
        response: ({ body }) => {
            const { title, content, tag } = body;
            const newData = {
                id: data.length + 1,
                title: title,
                content: content,
                tag: tag
            }

            if (!title || !content || !tag) {
                return {
                    code: 401,
                    data: { title: title, content: content, tag: tag },
                    message: '创建失败'
                }
            } else {
                return {
                    code: 200,
                    data: newData,
                    message: '创建成功'
                }
            }
        }
    },
    {
        url: '/api/deleteTopic',
        method: 'post',
        response: ({ body }) => {
            const id = body.id
            let filteredList = [...data]

            filteredList = filteredList.filter((item) => item.id !== id);

            return {
                code: 200,
                data: filteredList,
                message: '编辑成功'
            }
        }
    }
]