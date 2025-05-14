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
---------------以上是示例，全量接口待补充--------------