# XbookLab 系统设计与实现

## 一、系统设计
### 1.1 架构设计
系统采用分层架构，分为表现层、业务逻辑层和数据访问层：
- 表现层：使用Gin框架处理HTTP请求，定义API路由（如`main.go`中`/api/auth`、`/api/user`等路由分组）
- 业务逻辑层：通过`handlers`包处理具体业务逻辑（如`handlers.Login`处理登录认证）
- 数据访问层：利用GORM操作MySQL数据库（如`models/user.go`定义用户模型），Redis用于缓存热点数据（如`main.go`中初始化Redis连接）

### 1.2 模块划分
根据`main.go`路由设计，系统分为以下核心模块：
- **认证模块**：包含登录（`/api/auth/login`）、登出（`/api/auth/logout`）等接口，使用JWT实现权限校验（`middleware/AuthMiddleware`）
- **用户模块**：支持注册（`/api/user/register`）、资料修改（`/api/user/update`）、关注关系管理（`/api/user/follow/:id`）
- **书籍模块**：提供书籍列表查询（`/api/book/`）、详情获取（`/api/book/:id`）、封面上传（`/api/book/upload`）等功能
- **推荐模块**：基于用户行为和标签偏好生成推荐（`/api/recommend`接口调用`handlers.GetRecommendedBooks`）
- **话题模块**：管理话题的创建、获取与更新，核心模型为`Topic`（包含`TopicID`、`Title`、`Content`等字段），提供`GetRandomTopics`接口用于随机获取话题，与用户模块关联（通过`AuthorUserID`记录发布者）
- **评论模块**：处理评论的发布、查询与删除，核心模型为`Comment`（支持书籍/话题评论区分，通过`Type`字段标识），提供`CreateComment`（创建评论）、`GetComments`（获取评论列表）、`DeleteComment`（删除评论）接口，与用户模块（`UserID`关联评论者）、书籍/话题模块（`TargetID`关联被评论对象）交互
- **管理模块**：支持管理员对话题/评论的审核与管理，基于现有接口扩展权限控制（如管理员可调用`DeleteComment`删除违规内容），保障社区内容合规性

### 1.3 数据库设计
根据`sql/xbooklab.sql`表结构，核心数据表及关系如下：
| 表名         | 核心字段                  | 关联关系                                                                 |
|--------------|---------------------------|--------------------------------------------------------------------------|
| users        | user_id, username, avatar | 一对多：用户-评论（comments.user_id）；用户-评分（ratings.user_id）       |
| books        | book_id, title, author    | 一对多：书籍-评论（comments.target_id）；书籍-评分（ratings.book_id）     |
| recommend    | user_id, book_id, score   | 多对多：用户-推荐书籍（联合主键(user_id, book_id)）                      |
| user_tags    | user_id, tag_id           | 多对多：用户-兴趣标签（外键关联users.user_id和tags.tag_id）              |
| book_tags    | book_id, tag_id           | 多对多：书籍-内容标签（外键关联books.book_id和tags.tag_id）              |

## 二、系统实现
### 2.1 技术选型
- **Web框架**：Gin（轻量高效，支持中间件机制，如`middleware.CorsMiddleware`处理跨域）
- **ORM工具**：GORM（简化数据库操作，`models`包中所有模型均通过GORM定义）
- **缓存数据库**：Redis（用于存储会话信息，`main.go`中通过`initRedis`初始化连接）
- **关系数据库**：MySQL（存储结构化数据，`initMySQL`函数实现数据库初始化和连接）

### 2.2 关键功能实现
#### 2.2.1 认证功能
- 登录验证：`handlers.Login`校验用户密码后生成JWT令牌（存储于Redis，设置过期时间）
- 权限控制：`middleware.AuthMiddleware`解析请求头中的JWT，验证有效性后传递用户信息到后续处理

#### 2.2.2 推荐功能
- 算法实现：基于用户评分（`models/ratings`表）、浏览记录（`models/book_views`表）和关注关系（`models/follows`表）计算推荐分数
- 接口实现：`handlers.GetRecommendedBooks`调用`models.GetRecommendedBooks`，通过JOIN操作关联`recommend`表和`books`表获取推荐结果（代码见`models/recommend.go`）

#### 2.2.3 用户行为记录
- 浏览记录：`handlers.RecordBookView`将用户ID和书籍ID写入`book_views`表（`sql/xbooklab.sql`中定义）
- 评分记录：`handlers.RecordBookRating`校验评分有效性后更新`ratings`表，并同步更新`books`表的`average_rating`字段