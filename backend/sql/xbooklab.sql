-- 创建数据库
CREATE DATABASE IF NOT EXISTS x_book_lab_test DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE x_book_lab_test;
-- 用户表
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `username` VARCHAR(50) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `avatar` VARCHAR(255),
  `bio` TEXT,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 标签表
CREATE TABLE IF NOT EXISTS `tags` (
  `tag_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `tag_name` VARCHAR(50) UNIQUE NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户兴趣标签关联表
CREATE TABLE IF NOT EXISTS `user_tags` (
  `user_id` INT UNSIGNED,
  `tag_id` INT UNSIGNED,
  PRIMARY KEY (`user_id`, `tag_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE,
  FOREIGN KEY (`tag_id`) REFERENCES `tag`(`tag_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 书籍表
CREATE TABLE IF NOT EXISTS `books` (
  `book_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `author` VARCHAR(100),
  `cover` VARCHAR(255),
  `description` TEXT,
  `average_rating` DECIMAL(3,1) DEFAULT 0.0,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 书籍标签关联表
CREATE TABLE IF NOT EXISTS `book_tags` (
  `book_id` INT UNSIGNED,
  `tag_id` INT UNSIGNED,
  PRIMARY KEY (`book_id`, `tag_id`),
  FOREIGN KEY (`book_id`) REFERENCES `book`(`book_id`) ON DELETE CASCADE,
  FOREIGN KEY (`tag_id`) REFERENCES `tag`(`tag_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 话题表
CREATE TABLE IF NOT EXISTS `topics` (
  `topic_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `content` TEXT,
  `author_user_id` INT UNSIGNED,
  `like_count` INT UNSIGNED DEFAULT 0,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`author_user_id`) REFERENCES `user`(`user_id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 话题标签关联表
CREATE TABLE IF NOT EXISTS `topic_tags` (
  `topic_id` INT UNSIGNED,
  `tag_id` INT UNSIGNED,
  PRIMARY KEY (`topic_id`, `tag_id`),
  FOREIGN KEY (`topic_id`) REFERENCES `topic`(`topic_id`) ON DELETE CASCADE,
  FOREIGN KEY (`tag_id`) REFERENCES `tag`(`tag_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 评论表
CREATE TABLE IF NOT EXISTS `comments` (
  `comment_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `target_id` INT UNSIGNED NOT NULL, -- 对应书籍/话题ID
  `type` TINYINT UNSIGNED NOT NULL, -- 0:书籍评论 1:话题评论
  `content` TEXT NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 评分表
CREATE TABLE IF NOT EXISTS `rating` (
  `rating_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT UNSIGNED NOT NULL,
  `book_id` INT UNSIGNED NOT NULL,
  `rating_value` TINYINT UNSIGNED NOT NULL, -- 0-10分
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `user_book` (`user_id`, `book_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE,
  FOREIGN KEY (`book_id`) REFERENCES `book`(`book_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 浏览记录表
CREATE TABLE IF NOT EXISTS `book_view` (
  `view_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT UNSIGNED,
  `book_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE SET NULL,
  FOREIGN KEY (`book_id`) REFERENCES `book`(`book_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 关注表
CREATE TABLE IF NOT EXISTS `follows` (
  `follower_id` INT UNSIGNED NOT NULL,
  `followed_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`follower_id`, `followed_id`),
  FOREIGN KEY (`follower_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE,
  FOREIGN KEY (`followed_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 话题点赞表
CREATE TABLE IF NOT EXISTS `topic_like` (
  `like_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT UNSIGNED NOT NULL,
  `topic_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `user_topic` (`user_id`, `topic_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE,
  FOREIGN KEY (`topic_id`) REFERENCES `topic`(`topic_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 事件表（用户主页动态）
CREATE TABLE IF NOT EXISTS `event` (
  `event_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT UNSIGNED NOT NULL,
  `event_type` TINYINT UNSIGNED NOT NULL, -- 0:阅读 1:评论 2:评分 3:点赞话题
  `event_content` TEXT,
  `target_id` INT UNSIGNED, -- 书籍/话题ID
  `target_type` TINYINT UNSIGNED, -- 0:书籍 1:话题
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `user`(`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;