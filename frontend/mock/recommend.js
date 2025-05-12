import { message } from 'ant-design-vue'
import Mock from 'mockjs'

const { data } = Mock.mock({
    'data|20': [{
        'id|+1': 1,
        title: '@ctitle(3, 8)',
        author: '@cname',
        cover: () => Mock.Random.image('150x200', '#4A7BF7', '#FFF', 'png', '书'),
        rating: () => Mock.Random.float(5, 10, 1, 1),
        desc: '@cparagraph(2, 4)', // 2-4段中文描述
        'tags|1-3': ['@cword(2,4)'], // 1-3个中文标签
        // category: '@pick(["literature","technology","history","psychology","art","business","philosophy"])',
        tag: '@pick(["文学","科技","历史","心理学","艺术","商业","哲学"])',
    }]
})

export default [
    {
        url: '/api/recommend',
        method: 'get',
        response: ({ query }) => {

            const tag = query.tag || ''
            let filteredList = [...data]


            if (tag) {
                filteredList = data.filter(
                    item => item.tag === tag
                )
            }

            return {
                code: 200,
                data: filteredList,
                total: filteredList.length
            }
        }
    },
    {
        // 书籍详情
        url: '/api/getBookDetail/:id', // 动态路由参数
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
        // 登录状态下记录浏览书籍详情行为
        url: '/api/recordBookView',
        method: 'post',
        response: () => {
            return {
                code: 200,
                data: '',
                message: '阅读次数+1'
            }
        }
    },
    {
        //评论列表
        url: '/api/book/comments',
        method: 'get',
        response: () => {
            const { data } = Mock.mock({
                'data|5-10': [{
                    'id|+1': 1,
                    author: '@cname', // 随机中文名
                    avatar: () => `https://i.pravatar.cc/48?img=${Mock.Random.integer(1, 70)}`, // 1-70号头像
                    content: '@cparagraph(1, 3)', // 1-3段随机评论
                    time: () => {
                        const units = ['秒', '分钟', '小时', '天', '周', '月']
                        const val = Mock.Random.integer(1, 30)
                        return `${val}${units[Mock.Random.integer(0, 5)]}前` // 如：3小时前
                    },
                    'likes|0-100': 1, // 随机点赞数
                    'replies|0-3': [{ // 0-3条回复
                        'id|+1': 1000,
                        author: '@cname',
                        content: '@csentence(10, 30)',
                        time: '@now("yyyy-MM-dd HH:mm")'
                    }]
                }]
            })

            return {
                code: 200,
                data: data.sort((a, b) => b.id - a.id), // 按ID倒序
                total: data.length
            }
        }
    },

    {
        // 记录用户评论
        url: '/api/recordBookComment',
        method: 'post',
        response: ({ body }) => {
            const content = body.content.trim()
            const bookId = body.book_id;

            if (!content) {
                return {
                    code: 401,
                    data: '',
                    message: '评论不能为空'
                }
            } else {
                return {
                    code: 200,
                    data: '',
                    message: '评论成功'
                }
            }
        }
    },

    {
        // 记录用户评分
        url: '/api/recordBookRating',
        method: 'post',
        response: ({ body }) => {
            const bookId = body.book_id;
            const rate = body.rating;

            return {
                code: 200,
                data: '',
                message: '评分成功'
            }
        }
    },

    // 后台接口
    {
        url: '/api/manage/getBooks',
        method: 'get',
        response: () => {
            let filteredList = [...data]

            return {
                code: 200,
                data: filteredList,
                total: filteredList.length
            }
        }
    },
    {
        url: '/api/manage/createBook',
        method: 'post',
        response: ({ body }) => {
            const newBook = {
                id: data.length + 1,
                title: body.title,
                author: body.author,
                cover: body.cover,
                rating: body.rating,
                desc: body.desc,
                category: body.category,
            }

            return {
                code: 200,
                data: newBook,
                message: '创建成功'
            }
        }
    },
    {
        url: '/api/manage/editBook',
        method: 'post',
        response: ({ body }) => {
            const newBook = {
                id: body.id,
                title: body.title,
                author: body.author,
                cover: body.cover,
                rating: body.rating,
                desc: body.desc,
                category: body.category,
            }

            return {
                code: 200,
                data: newBook,
                message: '编辑成功'
            }
        }
    },
    {
        url: '/api/manage/delBook',
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
    },
    {
        url: '/api/manage/uploadImg',
        method: 'post',
        response: ({ body }) => {
            const feed_img = body.feed_img

            if (!feed_img) {
                return {
                    code: 200,
                    data: {
                        // 返回上传后的图片URL
                        url: 'https://bkimg.cdn.bcebos.com/pic/562c11dfa9ec8a136327f3867754868fa0ec08fa6a07?x-bce-process=image/format,f_auto/quality,Q_70/resize,m_lfit,limit_1,w_536'
                    },
                    message: '上传成功'
                }
            } else {
                return {
                    code: 401,
                    data: '',
                    message: '上传失败'
                }
            }


        }
    },
]