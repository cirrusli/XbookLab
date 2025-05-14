-- 插入标签
INSERT INTO tags (tag_name) VALUES 
('科技'), ('历史'), ('心理'), ('艺术'), ('哲学'), ('商业'), ('文学');

-- 插入用户, 密码为123456
INSERT INTO users (username, password, avatar, bio) VALUES
('ai_thinker', '$10$2uOLR.hS4Vtb1wUfhOyDYuabYSMu9ACVkz0T4aNjSTaw0OAQREsMW', 'https://randomuser.me/api/portraits/women/32.jpg', '热爱科幻小说与哲学思考'),
('history_man', '$10$2uOLR.hS4Vtb1wUfhOyDYuabYSMu9ACVkz0T4aNjSTaw0OAQREsMW', 'https://randomuser.me/api/portraits/men/33.jpg', '历史爱好者，喜欢读传记'),
('psychology_girl', '$10$2uOLR.hS4Vtb1wUfhOyDYuabYSMu9ACVkz0T4aNjSTaw0OAQREsMW', 'https://randomuser.me/api/portraits/women/34.jpg', '心理学研究生，研究行为模式'),
('artist', '$10$2uOLR.hS4Vtb1wUfhOyDYuabYSMu9ACVkz0T4aNjSTaw0OAQREsMW', 'https://randomuser.me/api/portraits/men/35.jpg', '艺术家，擅长油画创作'),
('elon-musk', '$10$2uOLR.hS4Vtb1wUfhOyDYuabYSMu9ACVkz0T4aNjSTaw0OAQREsMW', 'https://randomuser.me/api/portraits/men/36.jpg', '科技公司CEO，关注AI发展');

-- 用户兴趣标签关联数据
INSERT INTO user_tags (user_id, tag_id)
SELECT u.user_id, t.tag_id
FROM users u
CROSS JOIN tags t
WHERE
    (u.username = 'ai_thinker' AND t.tag_name IN ('科技', '哲学')) OR
    (u.username = 'history_man' AND t.tag_name IN ('历史', '艺术')) OR
    (u.username = 'psychology_girl' AND t.tag_name IN ('心理', '哲学')) OR
    (u.username = 'artist' AND t.tag_name = '艺术') OR
    (u.username = 'elon-musk' AND t.tag_name = '科技');

-- 插入书籍
INSERT INTO books (title, author, cover, description,average_rating) VALUES
('三体', '刘慈欣', 'https://img9.doubanio.com/view/subject/s/public/s2768378.jpg', '中国科幻文学的里程碑之作，讲述地球文明与三体文明的接触与冲突。', 9.0),
('人类简史', '尤瓦尔·赫拉利', 'https://img9.doubanio.com/view/subject/s/public/s27814883.jpg', '从动物到上帝的历史叙事，从认知革命到科学革命，重新审视人类历史的发展脉络。', 8.5),
('百年孤独', '加西亚·马尔克斯', 'https://img9.doubanio.com/view/subject/s/public/s6384944.jpg', '魔幻现实主义文学的代表作，讲述布恩迪亚家族七代人的传奇故事。', 8.0),
('黑镜：未来预言', '查理·布鲁克', 'https://img3.doubanio.com/view/subject/l/public/s33491253.jpg', '科技反乌托邦短篇集', 8.5),
('禅与摩托车维修艺术', '罗伯特·波西格', 'https://img3.doubanio.com/view/subject/l/public/s29831183.jpg', '哲学与旅行的融合', 8.0),
('活着', '余华', 'https://img9.doubanio.com/view/subject/s/public/s29053580.jpg', '讲述一个人在中国社会变革中的苦难与坚韧。',9.1),
('原则', '瑞·达利欧', 'https://img9.doubanio.com/view/subject/s/public/s29651109.jpg', '桥水基金创始人分享他的生活和工作原则。', 8.0),
('小王子', '圣埃克苏佩里', 'https://img9.doubanio.com/view/subject/s/public/s1103152.jpg', '一部写给大人的童话，探讨爱与责任的真谛。',9.4),
('思考，快与慢', '丹尼尔·卡尼曼', 'https://img9.doubanio.com/view/subject/s/public/s6907698.jpg', '诺贝尔经济学奖得主揭示人类思维的两种系统。',8.7),
('沉默的大多数', '王小波', 'https://img9.doubanio.com/view/subject/s/public/s1447349.jpg', '王小波杂文代表作，展现其独特的思想与文风。',8.2),
('局外人', '阿尔贝·加缪', 'https://img9.doubanio.com/view/subject/s/public/s4468484.jpg', '存在主义文学的代表作，讲述一个局外人的故事。',8.9),
('未来简史', '尤瓦尔·赫拉利', 'https://img9.doubanio.com/view/subject/s/public/s29287103.jpg', '探讨人类未来可能面临的重大变革与挑战。',7.8);

-- 书籍标签关联数据
INSERT INTO book_tags (book_id, tag_id)
SELECT b.book_id, t.tag_id
FROM books b
CROSS JOIN tags t
WHERE
    (b.title = '三体' AND t.tag_name = '科技') OR
    (b.title = '人类简史' AND t.tag_name = '历史') OR
    (b.title = '百年孤独' AND t.tag_name = '艺术') OR
    (b.title = '黑镜：未来预言' AND t.tag_name = '科技') OR
    (b.title = '禅与摩托车维修艺术' AND t.tag_name = '哲学') OR
    (b.title = '活着' AND t.tag_name = '历史') OR
    (b.title = '原则' AND t.tag_name = '心理') OR
    (b.title = '小王子' AND t.tag_name = '艺术') OR
    (b.title = '思考，快与慢' AND t.tag_name = '心理') OR
    (b.title = '沉默的大多数' AND t.tag_name = '哲学') OR
    (b.title = '局外人' AND t.tag_name = '哲学') OR
    (b.title = '未来简史' AND t.tag_name = '历史');

-- 插入话题
INSERT INTO topics (title, content, author_user_id,like_count) VALUES
('AI会统治人类吗？', '讨论AI伦理与未来社会的影响', 5, 1235),
('如何理解存在主义？', '从萨特到加缪的哲学探讨', 1, 3462),
('油画基础教程', '零基础学习油画技巧', 4, 869),
('历史小说推荐', '分享你读过的历史题材小说', 2, 934),
('冥想与心理健康', '心理学视角下的冥想实践', 3, 1212),
('分享你最喜欢的一本书', '一千个读者心中有一千个哈姆雷特', 6, 660);

-- 话题标签关联
INSERT INTO topic_tag (topic_id, tag_id) VALUES
(1, 1), (1, 5),  -- AI话题关联 科技和哲学
(2, 5),          -- 存在主义话题关联 哲学
(3, 4),          -- 油画话题关联 艺术
(4, 2),          -- 历史小说话题关联 历史
(5, 3);          -- 冥想话题关联 心理

-- 插入评分记录
INSERT INTO rating (user_id, book_id, rating_value) VALUES
(1, 1, 9),   -- ai_thinker 给《三体》9分
(1, 2, 8),   -- ai_thinker 给《人类简史》8分
(2, 2, 10),  -- history_man 给《人类简史》10分
(3, 5, 7),   -- psychology_girl 给《禅与摩托车维修艺术》7分
(4, 3, 8),   -- artist 给《百年孤独》8分
(5, 4, 9);   -- elon-musk 给《黑镜：未来预言》9分

-- 插入浏览记录
INSERT INTO book_view (user_id, book_id) VALUES
(1, 2), (1, 3),  -- ai_thinker 浏览《人类简史》《百年孤独》
(2, 1), (2, 2),  -- history_man 浏览《三体》《人类简史》
(3, 5),         -- psychology_girl 浏览《禅与摩托车维修艺术》
(4, 3),         -- artist 浏览《百年孤独》
(5, 4);         -- elon-musk 浏览《黑镜：未来预言》

-- 插入评论（区分书籍/话题）
INSERT INTO comment (target_id, type, content, user_id) VALUES
(1, 0, '书友小李说：《三体》让我重新思考宇宙的意义', 2),  -- 书籍评论
(2, 0, '书友小王说：历史与哲学的完美结合', 3),            -- 书籍评论
(1, 1, '话题用户A说：AI应该服务于人类而非统治', 4),      -- 话题评论
(2, 1, '话题用户B说：存在主义的核心是自由选择', 5);       -- 话题评论

-- 插入关注关系
INSERT INTO follow (follower_id, followed_id) VALUES
(1, 2), (1, 5),  -- ai_thinker 关注 history_man 和 elon-musk
(2, 1), (2, 3),  -- history_man 关注 ai_thinker 和 psychology_girl
(3, 4),         -- psychology_girl 关注 artist
(4, 1),         -- artist 关注 ai_thinker
(5, 2);         -- elon-musk 关注 history_man

-- 插入话题点赞
INSERT INTO topic_like (user_id, topic_id) VALUES
(1, 1), (2, 1),  -- ai_thinker 和 history_man 点赞话题《AI会统治人类吗？》
(3, 2),         -- psychology_girl 点赞话题《如何理解存在主义？》
(4, 3),         -- artist 点赞话题《油画基础教程》
(5, 5);         -- elon-musk 点赞话题《冥想与心理健康》

-- 插入事件记录（动态）
INSERT INTO event (user_id, event_type, event_content, target_id, target_type) VALUES
(1, 0, '阅读了《人类简史》', 2, 0),      -- 浏览事件
(1, 2, '评分《人类简史》为10分', 2, 0), -- 评分事件
(3, 1, '评论了《禅与摩托车维修艺术》', 5, 0), -- 评论事件
(4, 3, '点赞了话题《油画基础教程》', 3, 1),    -- 点赞事件
(5, 0, '阅读了《黑镜：未来预言》', 4, 0);      -- 浏览事件