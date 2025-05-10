-- 推荐表定义
CREATE TABLE IF NOT EXISTS recommend (
    user_id INT UNSIGNED NOT NULL COMMENT '用户ID',
    book_id INT UNSIGNED NOT NULL COMMENT '书籍ID',
    score FLOAT NOT NULL COMMENT '推荐分数',
    update_at DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (user_id, book_id),
    INDEX idx_user_id (user_id),
    INDEX idx_book_id (book_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户书籍推荐表';