# Golang Web 应用

这是一个使用 Golang 和 Gin 框架开发的 Web 应用，提供 JWT 认证功能和用户管理。

## 功能特性

- 🔐 JWT 认证系统
- 👤 用户登录和注册
- 🗄️ MySQL 数据库支持
- ⚙️ 环境配置管理
- 🛡️ 中间件认证保护
- 🔄 令牌刷新功能

## 技术栈

- **语言**: Go 1.21+
- **Web框架**: Gin
- **数据库**: MySQL
- **认证**: JWT (JSON Web Token)
- **配置管理**: Viper

## 项目结构

```
golang-web/
├── config/                 # 配置文件
│   ├── config.go          # 配置结构定义
│   ├── config.development.yaml  # 开发环境配置
│   └── config.production.yaml   # 生产环境配置
├── database/              # 数据库相关
│   └── database.go        # 数据库连接和初始化
├── models/                # 数据模型
│   └── user.go           # 用户模型
├── handlers/              # 请求处理器
│   └── auth.go           # 认证处理器
├── middleware/            # 中间件
│   └── auth.go           # JWT认证中间件
├── routes/                # 路由配置
│   └── routes.go         # 路由设置
├── utils/                 # 工具函数
│   └── jwt.go            # JWT工具
├── go.mod                 # Go模块文件
├── main.go                # 主程序
└── README.md              # 项目说明
```

## 环境要求

- Go 1.21 或更高版本
- MySQL 5.7 或更高版本
- 支持的环境变量: `GO_ENV`

## 安装和运行

### 1. 克隆项目

```bash
git clone <repository-url>
cd golang-web
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

#### 开发环境
- 数据库名: `golang_dev`
- 用户名: `root`
- 密码: `123456`
- 主机: `localhost`
- 端口: `3306`

#### 生产环境
- 数据库名: `golang_web`
- 用户名: `root`
- 密码: `123456`
- 主机: `localhost`
- 端口: `3306`

### 4. 设置环境变量

```bash
# 开发环境
export GO_ENV=development

# 生产环境
export GO_ENV=production
```

### 5. 运行应用

```bash
go run main.go
```

应用将在 `http://localhost:8080` 启动。

## API 接口

### 认证接口

#### 用户登录
```
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

#### 用户注册
```
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "newuser",
  "password": "password123",
  "email": "user@example.com"
}
```

### 受保护的接口

#### 获取用户信息
```
GET /api/v1/user/profile
Authorization: Bearer <jwt_token>
```

#### 刷新令牌
```
POST /api/v1/token/refresh
Authorization: Bearer <jwt_token>
```

### 健康检查
```
GET /health
```

## 默认用户

应用启动时会自动创建默认用户：
- 用户名: `admin`
- 密码: `admin123`

## 配置说明

### 开发环境配置 (`config.development.yaml`)
- 服务器模式: `debug`
- 端口: `8080`
- JWT密钥: `dev-secret-key-change-in-production`
- JWT过期时间: `24` 小时

### 生产环境配置 (`config.production.yaml`)
- 服务器模式: `release`
- 端口: `8080`
- JWT密钥: `your-secret-key-change-in-production`
- JWT过期时间: `24` 小时

## 安全注意事项

1. **生产环境**: 请修改默认的 JWT 密钥
2. **密码安全**: 当前使用简单密码比较，生产环境建议使用 bcrypt
3. **数据库安全**: 请修改默认数据库密码
4. **HTTPS**: 生产环境建议启用 HTTPS

## 开发说明

### 添加新的API接口

1. 在 `handlers/` 目录下创建新的处理器
2. 在 `routes/routes.go` 中添加路由
3. 根据需要添加中间件

### 数据库迁移

当前使用简单的表创建语句，生产环境建议使用专业的数据库迁移工具。

### 日志

应用使用标准库的 `log` 包，生产环境建议使用结构化日志库。

## 故障排除

### 数据库连接失败
- 检查 MySQL 服务是否运行
- 验证数据库连接参数
- 确认数据库用户权限

### JWT 验证失败
- 检查令牌格式是否正确
- 验证令牌是否过期
- 确认 JWT 密钥配置

## 许可证

本项目采用 MIT 许可证。
