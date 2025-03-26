# API 文档

## 用户认证

### 登录
`POST /api/auth/login`

请求参数:
```json
{
  "username": "string",
  "password": "string"
}
```

响应:
```json
{
  "token": "string",
  "user": {
    "id": "number",
    "username": "string"
  }
}
```

### 登出
`POST /api/auth/logout`

响应:
```json
{
  "message": "登出成功"
}
```

## 图书相关

### 获取推荐图书列表
`GET /api/books/recommend`

响应:
```json
[
  {
    "id": "number",
    "title": "string",
    "cover": "string",
    "rating": "number",
    "description": "string",
    "category": "string"
  }
]
```

### 获取图书详情
`GET /api/books/:id`

响应:
```json
{
  "id": "number",
  "title": "string",
  "cover": "string",
  "author": "string",
  "publisher": "string",
  "rating": "number",
  "description": "string",
  "category": "string"
}
```

### 获取分类列表
`GET /api/categories`

响应:
```json
[
  {
    "id": "number",
    "name": "string"
  }
]
```

## 用户相关

### 修改密码
`POST /api/auth/change-password`

请求参数:
```json
{
  "oldPassword": "string",
  "newPassword": "string"
}
```

响应:
```json
{
  "message": "密码修改成功"
}
```

### 注册
`POST /api/auth/register`

请求参数:
```json
{
  "username": "string",
  "password": "string"
}
```

响应:
```json
{
  "message": "注册成功"
}
```