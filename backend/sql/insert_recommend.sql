-- 批量推荐数据插入语句

INSERT INTO `recommend` (`user_id`, `book_id`, `score`, `update_at`) VALUES (1, 2, 9.1, NOW()), (1, 2, 9.6, NOW()), (1, 48, 8.1, NOW()), (1, 3, 9.6, NOW()), (1, 48, 7.5, NOW()), (1, 9, 9.9, NOW()), (1, 12, 9.3, NOW()), (1, 21, 9.8, NOW()), (1, 46, 8.2, NOW()), (1, 7, 9.4, NOW()), (1, 29, 9.3, NOW()), (1, 21, 9.6, NOW()), (1, 30, 9.5, NOW()), (1, 41, 9.0, NOW()), (1, 51, 7.9, NOW()), (1, 38, 8.9, NOW()), (1, 19, 9.5, NOW()), (1, 21, 9.3, NOW()), (1, 31, 9.9, NOW()), (1, 1, 10.0, NOW());

-- user_id=6
INSERT INTO recommend (user_id, book_id, score, update_at)
VALUES
(6, 1, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 2, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 3, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 4, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 5, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 6, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 7, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 8, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 9, ROUND(7 + RAND() * 3, 2), NOW()),
(6, 10, ROUND(7 + RAND() * 3, 2), NOW());