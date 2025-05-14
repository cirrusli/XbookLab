# XbookLab 推荐算法技术文档

## 一、数据库表结构说明

### 相关数据表及核心字段
| 表名          | 字段名          | 类型       | 说明                     |
|---------------|-----------------|------------|--------------------------|
| users         | user_id         | uint        | 用户ID（主键）           |
| ratings       | user_id         | uint        | 用户ID（外键）           |
|               | book_id         | uint        | 书籍ID（外键）           |
|               | rating_value    | float       | 用户对书籍的评分（1-10分）|
| book_views    | user_id         | uint        | 用户ID（外键）           |
|               | book_id         | uint        | 书籍ID（外键）           |
| user_tags     | user_id         | uint        | 用户ID（外键）           |
|               | tag_id          | uint        | 标签ID（外键）           |
| follow        | follower_id     | uint        | 关注者ID（外键）         |
|               | followed_id     | uint        | 被关注者ID（外键）       |
| book_tags     | book_id         | uint        | 书籍ID（外键）           |
|               | tag_id          | uint        | 标签ID（外键）           |
| recommend     | user_id         | uint        | 用户ID（外键）           |
|               | book_id         | uint        | 书籍ID（外键）           |
|               | score           | float       | 推荐分数（越高越相关）   |
|               | update_at       | datetime    | 推荐结果更新时间         |

## 二、代码执行链路

### 1. 用户行为触发推荐评分更新
当用户产生评分、浏览等行为时，通过 `UpsertRecommendScore` 函数更新推荐分数（`recommend` 表）：
```go
func UpsertRecommendScore(userID uint, bookID uint, score float64) error {
    return DB.Transaction(func(tx *gorm.DB) error {
        err := tx.Exec(`
            INSERT INTO recommend (user_id, book_id, score, update_at)
            VALUES (?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE score = score + ?, update_at = ?
        `, userID, bookID, score, time.Now(), score, time.Now()).Error
        return err
    })
}
```
该函数通过事务实现「插入或更新」逻辑：若记录已存在则累加分数，否则新增记录。

### 2. 协同过滤算法计算
核心逻辑位于 `tools/import_recommend.go`，执行流程如下：

#### (1) 数据采集
- 从 `users` 表获取所有用户ID
- 对每个用户：
  - 从 `ratings` 获取评分记录（`getUserRatings`）
  - 从 `book_views` 获取浏览记录（`getUserViews`）
  - 从 `user_tags` 获取兴趣标签（`getUserTags`）
  - 从 `follow` 获取关注用户（`getFollowedUsers`）

#### (2) 相似用户计算（`calculateSimilarUsers`）
- **标签相似度**：使用Jaccard系数计算用户兴趣标签的交集/并集比例（`calculateJaccardSimilarity`）
- **评分相似度**：使用余弦相似度计算用户评分向量的相似性（`calculateCosineSimilarity`）
- **综合相似度**：标签相似度×0.3 + 评分相似度×0.7，筛选相似度>0.3的前10名用户

#### (3) 推荐书籍生成
- **来源1（相似用户/关注用户）**：提取来源用户评分≥7的书籍（排除用户已评分/浏览的书籍）
- **来源2（兴趣标签）**：从 `book_tags` 提取标签匹配度高的书籍（排除用户已交互书籍）
- **来源3（探索性推荐）**：选取用户未接触但全局评分≥8的书籍（标签与用户兴趣无交集）

### 3. 推荐结果存储
通过 `insertRecommendations` 函数批量插入 `recommend` 表：
```go
func insertRecommendations(db *sql.DB, recommendations []Recommendation) {
    tx, err := db.Begin()
    if err != nil { log.Fatal(err) }
    // 清空旧数据
    _, err = tx.Exec("TRUNCATE TABLE recommend")
    // 批量插入新数据
    stmt, _ := tx.Prepare("INSERT INTO recommend(user_id, book_id, score, update_at) VALUES(?, ?, ?, ?)")
    for _, rec := range recommendations { stmt.Exec(rec.UserID, rec.BookID, rec.Score, rec.Updated) }
    tx.Commit()
}
```

## 三、关键算法说明
- **Jaccard相似度**：用于衡量用户兴趣标签的重叠程度，公式：`|A∩B| / |A∪B|`
- **余弦相似度**：用于衡量用户评分向量的方向相似性，公式：`(A·B) / (|A|×|B|)`
- **推荐分数加权**：相似用户评分（权重0.6）+ 标签匹配度（权重0.4）+ 探索性评分（权重0.3）

## 四、数据表详细说明

### 1. recommend（用户书籍推荐表）
- **功能**：存储用户与书籍的推荐关系及推荐分数
- **字段解释**：
  | 字段名     | 类型            | 说明               |
  |------------|-----------------|--------------------|
  | user_id    | INT UNSIGNED    | 用户ID（主键）     |
  | book_id    | INT UNSIGNED    | 书籍ID（主键）     |
  | score      | FLOAT           | 推荐分数（越高越相关） |
  | update_at  | DATETIME        | 推荐结果更新时间   |

### 2. users（用户表）
- **功能**：存储用户基础信息
- **字段解释**：
  | 字段名     | 类型            | 说明               |
  |------------|-----------------|--------------------|
  | user_id    | INT UNSIGNED    | 用户ID（自增主键） |
  | username   | VARCHAR(50)     | 用户名（唯一）     |
  | password   | VARCHAR(255)    | 密码               |
  | avatar     | VARCHAR(255)    | 头像地址           |
  | bio        | TEXT            | 个人简介           |
  | created_at | TIMESTAMP       | 创建时间           |

### 3. tags（标签表）
- **功能**：存储通用标签信息
- **字段解释**：
  | 字段名     | 类型            | 说明               |
  |------------|-----------------|--------------------|
  | tag_id     | INT UNSIGNED    | 标签ID（自增主键） |
  | tag_name   | VARCHAR(50)     | 标签名（唯一）     |
  | created_at | TIMESTAMP       | 创建时间           |

### 4. user_tags（用户兴趣标签关联表）
- **功能**：存储用户与兴趣标签的关联关系
- **字段解释**：
  | 字段名     | 类型            | 说明               |
  |------------|-----------------|--------------------|
  | user_id    | INT UNSIGNED    | 用户ID（外键）     |
  | tag_id     | INT UNSIGNED    | 标签ID（外键）     |
  | created_at | TIMESTAMP       | 创建时间           |

### 5. books（书籍表）
- **功能**：存储书籍基础信息
- **字段解释**：
  | 字段名          | 类型            | 说明               |
  |-----------------|-----------------|--------------------|
  | book_id         | INT UNSIGNED    | 书籍ID（自增主键） |
  | title           | VARCHAR(255)    | 书名（非空）       |
  | author          | VARCHAR(100)    | 作者               |
  | cover           | VARCHAR(255)    | 封面地址           |
  | description     | TEXT            | 书籍描述           |
  | average_rating  | DECIMAL(3,1)    | 平均评分（默认0.0） |
  | created_at      | TIMESTAMP       | 创建时间           |

### 6. book_tags（书籍标签关联表）
- **功能**：存储书籍与标签的关联关系
- **字段解释**：
  | 字段名     | 类型            | 说明               |
  |------------|-----------------|--------------------|
  | book_id    | INT UNSIGNED    | 书籍ID（外键）     |
  | tag_id     | INT UNSIGNED    | 标签ID（外键）     |

### 7. topics（话题表）
- **功能**：存储用户创建的讨论话题
- **字段解释**：
  | 字段名          | 类型            | 说明               |
  |-----------------|-----------------|--------------------|
  | topic_id        | INT UNSIGNED    | 话题ID（自增主键） |
  | title           | VARCHAR(255)    | 话题标题（非空）   |
  | content         | TEXT            | 话题内容           |
  | author_user_id  | INT UNSIGNED    | 作者用户ID（外键） |
  | like_count      | INT UNSIGNED    | 点赞数（默认0）    |
  | created_at      | TIMESTAMP       | 创建时间           |

### 8. topic_tags（话题标签关联表）
- **功能**：存储话题与标签的关联关系
- **字段解释**：
  | 字段名     | 类型            | 说明               |
  |------------|-----------------|--------------------|
  | topic_id   | INT UNSIGNED    | 话题ID（外键）     |
  | tag_id     | INT UNSIGNED    | 标签ID（外键）     |

### 9. comments（评论表）
- **功能**：存储书籍/话题的评论内容
- **字段解释**：
  | 字段名      | 类型            | 说明               |
  |-------------|-----------------|--------------------|
  | comment_id  | INT UNSIGNED    | 评论ID（自增主键） |
  | target_id   | INT UNSIGNED    | 目标ID（书籍/话题） |
  | type        | TINYINT UNSIGNED| 目标类型（0:书籍 1:话题） |
  | content     | TEXT            | 评论内容（非空）   |
  | user_id     | INT UNSIGNED    | 用户ID（外键）     |
  | created_at  | TIMESTAMP       | 创建时间           |

### 10. ratings（评分表）
- **功能**：存储用户对书籍的评分记录
- **字段解释**：
  | 字段名         | 类型            | 说明               |
  |----------------|-----------------|--------------------|
  | rating_id      | INT UNSIGNED    | 评分ID（自增主键） |
  | user_id        | INT UNSIGNED    | 用户ID（外键）     |
  | book_id        | INT UNSIGNED    | 书籍ID（外键）     |
  | rating_value   | TINYINT UNSIGNED| 评分值（0-10分）   |
  | created_at     | TIMESTAMP       | 创建时间           |
  | updated_at     | TIMESTAMP       | 最后更新时间       |

### 11. book_views（浏览记录表）
- **功能**：存储用户浏览书籍的记录
- **字段解释**：
  | 字段名      | 类型            | 说明               |
  |-------------|-----------------|--------------------|
  | view_id     | INT UNSIGNED    | 浏览记录ID（自增主键） |
  | user_id     | INT UNSIGNED    | 用户ID（外键）     |
  | book_id     | INT UNSIGNED    | 书籍ID（外键）     |
  | created_at  | TIMESTAMP       | 创建时间           |

### 12. follows（关注表）
- **功能**：存储用户之间的关注关系
- **字段解释**：
  | 字段名        | 类型            | 说明               |
  |---------------|-----------------|--------------------|
  | follower_id   | INT UNSIGNED    | 关注者ID（外键）   |
  | followed_id   | INT UNSIGNED    | 被关注者ID（外键） |
  | created_at    | TIMESTAMP       | 创建时间           |

### 13. topic_like（话题点赞表）
- **功能**：存储用户对话题的点赞记录
- **字段解释**：
  | 字段名      | 类型            | 说明               |
  |-------------|-----------------|--------------------|
  | like_id     | INT UNSIGNED    | 点赞ID（自增主键） |
  | user_id     | INT UNSIGNED    | 用户ID（外键）     |
  | topic_id    | INT UNSIGNED    | 话题ID（外键）     |
  | created_at  | TIMESTAMP       | 创建时间           |

### 14. event（用户主页动态表）
- **功能**：存储用户行为动态（如阅读、评论等）
- **字段解释**：
  | 字段名        | 类型            | 说明               |
  |---------------|-----------------|--------------------|
  | event_id      | INT UNSIGNED    | 动态ID（自增主键） |
  | user_id       | INT UNSIGNED    | 用户ID（外键）     |
  | event_type    | TINYINT UNSIGNED| 动态类型（0:阅读 1:评论 2:评分 3:点赞话题） |
  | event_content | TEXT            | 动态内容           |
  | target_id     | INT UNSIGNED    | 目标ID（书籍/话题） |
  | target_type   | TINYINT UNSIGNED| 目标类型（0:书籍 1:话题） |
  | created_at    | TIMESTAMP       | 创建时间           |
