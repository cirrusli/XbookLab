interface Book {
  id: number
  title: string
  cover: string
  author: string
  press: string
  tags: string[]
  collects: number
  discussions: number
}

export const mockBooks: Book[] = [
  {
    id: 1,
    title: 'JavaScript高级程序设计',
    cover: '/placeholder.jpg',
    author: '尼古拉斯·泽卡斯',
    press: '人民邮电出版社',
    tags: ['前端', '编程', 'JavaScript'],
    collects: 1422,
    discussions: 235
  },
  // 更多模拟数据...
]

localStorage.setItem('books', JSON.stringify(mockBooks))