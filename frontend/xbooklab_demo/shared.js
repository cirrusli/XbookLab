// 共享的数据和函数
const topics = [
    {
        id: 1,
        title: "如何培养深度阅读的习惯？",
        excerpt: "在这个信息爆炸的时代，我们越来越难以专注阅读一本书。分享一些我实践过的深度阅读方法...",
        likes: 245,
        category: "阅读方法"
    },
    {
        id: 2,
        title: "2023年最值得阅读的10本科技书籍",
        excerpt: "从人工智能到量子计算，这些书将带你了解科技前沿的发展趋势...",
        likes: 189,
        category: "科技"
    },
    {
        id: 3,
        title: "村上春树新作《城市与它的不确定的墙》读后感",
        excerpt: "时隔多年，村上春树再次带来一部充满隐喻与哲思的长篇小说...",
        likes: 320,
        category: "文学"
    },
    {
        id: 4,
        title: "心理学入门：从这5本书开始",
        excerpt: "想了解心理学却不知从何开始？这些入门书籍将为你打开心理学的大门...",
        likes: 156,
        category: "心理学"
    },
    {
        id: 5,
        title: "历史小说中的真实与虚构",
        excerpt: "探讨历史小说如何在尊重史实的基础上进行艺术创作...",
        likes: 98,
        category: "历史"
    },
    {
        id: 6,
        title: "科幻小说对未来科技的预测准确吗？",
        excerpt: "从《1984》到《神经漫游者》，科幻作品中的哪些预测已经成为现实...",
        likes: 210,
        category: "科技"
    }
];

const books = [
    {
        id: 1,
        title: "百年孤独",
        author: "加西亚·马尔克斯",
        cover: "https://img9.doubanio.com/view/subject/s/public/s6384944.jpg",
        rating: "9.3",
        desc: "魔幻现实主义文学的代表作，讲述布恩迪亚家族七代人的传奇故事。"
    },
    {
        id: 2,
        title: "人类简史",
        author: "尤瓦尔·赫拉利",
        cover: "https://img9.doubanio.com/view/subject/s/public/s27814883.jpg",
        rating: "9.1",
        desc: "从认知革命到科学革命，重新审视人类历史的发展脉络。"
    },
    {
        id: 3,
        title: "三体",
        author: "刘慈欣",
        cover: "https://img9.doubanio.com/view/subject/s/public/s2768378.jpg",
        rating: "8.8",
        desc: "中国科幻文学的里程碑之作，讲述地球文明与三体文明的接触与冲突。"
    },
    {
        id: 4,
        title: "活着",
        author: "余华",
        cover: "https://img9.doubanio.com/view/subject/s/public/s29053580.jpg",
        rating: "9.4",
        desc: "讲述一个人在中国社会变革中的苦难与坚韧。"
    },
    {
        id: 5,
        title: "原则",
        author: "瑞·达利欧",
        cover: "https://img9.doubanio.com/view/subject/s/public/s29651109.jpg",
        rating: "8.4",
        desc: "桥水基金创始人分享他的生活和工作原则。"
    },
    {
        id: 6,
        title: "小王子",
        author: "圣埃克苏佩里",
        cover: "https://img9.doubanio.com/view/subject/s/public/s1103152.jpg",
        rating: "9.0",
        desc: "一部写给大人的童话，探讨爱与责任的真谛。"
    },
    {
        id: 7,
        title: "思考，快与慢",
        author: "丹尼尔·卡尼曼",
        cover: "https://img9.doubanio.com/view/subject/s/public/s6907698.jpg",
        rating: "8.1",
        desc: "诺贝尔经济学奖得主揭示人类思维的两种系统。"
    },
    {
        id: 8,
        title: "沉默的大多数",
        author: "王小波",
        cover: "https://img9.doubanio.com/view/subject/s/public/s1447349.jpg",
        rating: "9.1",
        desc: "王小波杂文代表作，展现其独特的思想与文风。"
    },
    {
        id: 9,
        title: "局外人",
        author: "阿尔贝·加缪",
        cover: "https://img9.doubanio.com/view/subject/s/public/s4468484.jpg",
        rating: "9.0",
        desc: "存在主义文学的代表作，讲述一个局外人的故事。"
    },
    {
        id: 10,
        title: "未来简史",
        author: "尤瓦尔·赫拉利",
        cover: "https://img9.doubanio.com/view/subject/s/public/s29287103.jpg",
        rating: "8.4",
        desc: "探讨人类未来可能面临的重大变革与挑战。"
    }
];

// 共享的header加载逻辑
document.addEventListener('DOMContentLoaded', () => {
    if (document.getElementById('shared-header')) {
        fetch('header.html')
            .then(response => response.text())
            .then(html => {
                document.getElementById('shared-header').innerHTML = html;
                initHeader();
            });
    }
});

function initHeader() {
    const loginBtn = document.getElementById('login-btn');
    const logoutBtn = document.getElementById('logout-btn');
}

// 初始化共享header
document.addEventListener('DOMContentLoaded', function () {
    const header = document.getElementById('shared-header');
    if (header) {
        header.innerHTML = `
            <div class="nav-container">
                <ul class="nav-links">
                    <li><a href="./xbooklab.html">推荐</a></li>
                    <li><a href="./groups.html">兴趣小组</a></li>
                    <li><a href="./hot.html">热榜</a></li>
                    <li class="active"><a href="./friends.html">好友在读</a></li>
                    <li><a href="#" id="match-trigger">书友匹配</a></li>
                </ul>
                <div class="user-section">
                    <button class="login-btn" id="login-btn">登录</button>
                    <div class="user-profile" id="user-profile" style="display: none;">
                    <a href="/user" class="profile-link">
                        <img src="" alt="用户头像" class="user-avatar">
                        <span class="username"></span>
                          </a>
                    </div>
                </div>
            </div>
        `;

        // 高亮当前页面导航项
        const currentPage = window.location.pathname.split('/').pop() || './xbooklab.html';
        document.querySelectorAll('.nav-links li').forEach(li => {
            const link = li.querySelector('a');
            if (link && link.getAttribute('href') === currentPage) {
                li.classList.add('active');
            } else {
                li.classList.remove('active');
            }
        });
    }
});

// 用户数据模型
const currentUser = {
    id: 1,
    name: "书友小李",
    avatar: "https://randomuser.me/api/portraits/men/67.jpg",
    bio: "热爱阅读与分享，专注文学与哲学。通过阅读探索世界，在书籍中寻找智慧与共鸣。",
    stats: {
        following: 12,
        followers: 45,
        topics: 23,
        books: 156
    },
    reading: {
        id: 3,
        title: "三体",
        author: "刘慈欣",
        genre: "科幻",
        cover: "https://img9.doubanio.com/view/subject/s/public/s2768378.jpg"
    },
    following: [
        {
            id: 101,
            name: "黑格尔",
            avatar: "https://randomuser.me/api/portraits/men/32.jpg"
        },
        {
            id: 102,
            name: "王小波",
            avatar: "https://randomuser.me/api/portraits/women/44.jpg"
        },
        {
            id: 103,
            name: "张爱玲",
            avatar: "https://randomuser.me/api/portraits/women/22.jpg"
        }
    ],
    activities: [
        {
            type: "topic",
            content: "如何理解《存在与时间》中的此在概念",
            time: "2小时前"
        },
        {
            type: "reading",
            content: "《三体》 并添加了书评",
            time: "昨天"
        },
        {
            type: "follow",
            content: "黑格尔 并收藏了他的书单",
            time: "3天前"
        },
        {
            type: "group",
            content: "哲学与思辨",
            time: "5天前"
        },
        {
            type: "finished",
            content: "《人类简史》",
            time: "1周前"
        }
    ]
};

// 路由处理函数
function handleRoute(path) {
    switch (path) {
        case '/user':
            window.location.href = 'user.html';
            break;
        case '/':
        case '/xbooklab.html':
            window.location.href = 'xbooklab.html';
            break;
        case '/groups':
            window.location.href = 'groups.html';
            break;
        case '/hot':
            window.location.href = 'hot.html';
            break;
        case '/friends':
            window.location.href = 'friends.html';
            break;
    }
}

// 初始化路由事件
function initRouter() {
    document.querySelectorAll('a[href^="/"]').forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const path = new URL(link.href).pathname;
            handleRoute(path);
        });
    });
}

// 初始化用户主页
function initUserProfile() {
    if (window.location.pathname.includes('user.html')) {
        const user = currentUser;
        
        // 更新头部信息
        document.querySelector('.profile-name').textContent = user.name;
        document.querySelector('.profile-bio').textContent = user.bio;
        document.querySelector('.profile-avatar').src = user.avatar;
        
        // 更新统计数据
        document.querySelectorAll('.profile-stat strong')[0].textContent = user.stats.following;
        document.querySelectorAll('.profile-stat strong')[1].textContent = user.stats.followers;
        document.querySelectorAll('.profile-stat strong')[2].textContent = user.stats.topics;
        document.querySelectorAll('.profile-stat strong')[3].textContent = user.stats.books;
        
        // 更新最近在读
        const readingBook = user.reading;
        document.querySelector('.book-cover').src = readingBook.cover;
        document.querySelector('.book-info h4').textContent = readingBook.title;
        document.querySelector('.book-info p').textContent = `${readingBook.author} · ${readingBook.genre}`;
        
        // 更新动态
        const activityContainer = document.querySelector('.profile-activity');
        user.activities.forEach(activity => {
            const activityItem = document.createElement('div');
            activityItem.className = 'activity-item';
            activityItem.innerHTML = `
                <p>${getActivityText(activity)}</p>
                <p class="activity-time">${activity.time}</p>
            `;
            activityContainer.appendChild(activityItem);
        });
    }
}

function getActivityText(activity) {
    switch(activity.type) {
        case 'topic': return `发表了话题 <a href="#">"${activity.content}"</a>`;
        case 'reading': return `开始阅读 <a href="#">${activity.content}</a>`;
        case 'follow': return `关注了 <a href="#">${activity.content}</a>`;
        case 'group': return `加入了书友圈 <a href="#">"${activity.content}"</a>`;
        case 'finished': return `完成了 <a href="#">${activity.content}</a> 的阅读`;
        default: return activity.content;
    }
}

// 在DOM加载完成后初始化
document.addEventListener('DOMContentLoaded', () => {
    initHeader();
    initUserProfile();
});