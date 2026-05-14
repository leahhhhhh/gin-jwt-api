# gin-demo

基于 Gin 框架的 Go Web 项目脚手架，包含用户注册、登录和 JWT 鉴权功能。

## 技术栈

- Web 框架: [Gin](https://github.com/gin-gonic/gin)
- ORM: [GORM](https://gorm.io/)
- 数据库: MySQL
- JWT: [golang-jwt](https://github.com/golang-jwt/jwt)

## 项目结构

```
jw_go/
├── main.go           # 程序入口
├── go.mod            # Go 模块依赖
├── config/           # 配置相关
├── controller/       # 控制器层
├── middleware/       # 中间件
├── model/            # 数据模型
├── service/          # 业务逻辑层
├── router/           # 路由配置
└── utils/            # 工具函数
```

## 接口文档

| 方法 | 路径 | 说明 | 是否需要鉴权 |
|------|------|------|-------------|
| POST | /register | 用户注册 | 否 |
| POST | /login | 用户登录 | 否 |
| GET | /auth/info | 获取用户信息 | 是 |

### 1. 用户注册

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'
```

### 2. 用户登录

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'
```

返回示例：
```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 3. 获取用户信息（需要鉴权）

```bash
curl -X GET http://localhost:8080/auth/info \
  -H "Authorization: Bearer <your-token>"
```

## 快速开始

### 1. 配置数据库

修改数据库连接配置（DSN）。

### 2. 安装依赖

```bash
go mod download
```

### 3. 运行项目

```bash
go run main.go
```

服务默认运行在 `http://localhost:8080`

## JWT 说明

- Token 有效期：24 小时
- 需在请求头中携带：`Authorization: Bearer <token>`
