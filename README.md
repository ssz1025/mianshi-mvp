# Gin 项目最佳实践模板

这是一个基于 Gin 框架的 Go Web 应用最佳实践模板，采用领域驱动设计（DDD）风格的分层架构，包含了完整的项目结构和常用功能实现。

> 🚀 **快速开始**: 点击右上角 **"Use this template"** 按钮创建你的项目，或查看 [使用指南](.github/SETUP.md)

## 特性

### 核心功能

✨ **清晰的项目结构** - 采用 DDD 分层架构，职责明确，易于维护
🔐 **JWT 认证** - 完整的用户认证和授权实现
📝 **统一响应格式** - 标准化的 API 响应结构
✅ **参数验证** - 基于 validator 的请求参数验证
🔄 **中间件支持** - 日志、恢复、CORS、限流等常用中间件
💾 **数据库集成** - 使用 GORM 进行数据库操作
📊 **结构化日志** - 基于 zap 的结构化日志记录
🐳 **Docker 支持** - 包含 Dockerfile 和 docker-compose
🔥 **热重载** - 使用 Air 实现开发时热重载
🧪 **完整测试** - Repository 单元测试示例

### 高级功能

📚 **Swagger 文档** - 自动生成交互式 API 文档
🧹 **验证中间件** - 通用的 JSON 验证中间件
📈 **Pprof 分析** - 内置性能分析工具
🔍 **Sentry 监控** - 实时错误追踪和监控
🔗 **OpenTelemetry** - 分布式追踪支持

### 扩展功能

🔐 **Casbin 权限管理** - RBAC 访问控制，支持数据权限范围
☁️ **OSS 存储** - 阿里云 OSS 集成，支持 STS 临时凭证
💉 **Wire 依赖注入** - 编译时依赖注入，零运行时开销

### 开发工具

🧪 **REST Client** - VS Code 中直接测试 API
🎣 **Pre-commit Hooks** - 提交前自动代码检查
📏 **golangci-lint** - 全面的代码质量检查
⚙️ **EditorConfig** - 统一的编辑器配置
🤖 **GitHub Actions** - 自动化 CI/CD 流程

> 📖 **详细使用说明**: [架构说明](docs/ARCHITECTURE.md) | [功能指南](docs/FEATURES.md) | [开发工具](docs/DEV_TOOLS.md)

## 项目结构

```
gin-template/
├── cmd/server/               # 应用入口
├── internal/                 # 私有应用代码
│   ├── api/                  # API 层（handler/middleware/router）
│   ├── service/              # 业务逻辑层
│   ├── repository/           # 数据访问层
│   ├── model/                # 数据模型
│   ├── dto/                  # 数据传输对象
│   └── infra/                # 基础设施层
├── pkg/                      # 可复用的公共库（casbin/oss/jwt/...）
├── openapi/                  # Swagger/OpenAPI 生成物
├── docs/                     # 技术文档
├── config/                   # 配置文件
├── migrations/               # 数据库迁移
└── scripts/                  # 脚本工具
```

## 快速开始

### 前置要求

- Go 1.21+
- PostgreSQL 15+
- Redis 7+ (可选)

### 安装依赖

```bash
go mod tidy
```

### 配置数据库

1. 创建数据库：

```bash
createdb gin_template
```

或使用 Makefile：

```bash
make init-db
```

2. 修改配置文件 `config/config.yaml`：

```yaml
database:
  host: localhost
  port: 5432
  database: gin_template
  username: postgres
  password: your_password
```

### 运行应用

**直接运行：**

```bash
go run cmd/server/main.go
```

或使用 Makefile：

```bash
make run
```

**使用热重载（需要安装 Air）：**

```bash
# 安装 Air
go install github.com/cosmtrek/air@latest

# 运行
air
```

或：

```bash
make dev
```

**使用 Docker Compose：**

```bash
docker-compose up
```

应用将在 `http://localhost:8080` 启动。

### 访问 Swagger 文档

启动应用后访问：

```
http://localhost:8080/swagger/index.html
```

生成/更新 Swagger 文档：

```bash
make swagger
```

## API 文档

### 健康检查

```bash
GET /health
```

### 用户认证

**注册用户：**

```bash
POST /api/v1/users
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123",
  "age": 25
}
```

**登录：**

```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

返回示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user_id": "123",
    "username": "testuser"
  }
}
```

### 用户管理

**获取用户列表：**

```bash
GET /api/v1/users?page=1&page_size=10
```

**获取用户详情：**

```bash
GET /api/v1/users/:id
```

**更新用户（需要认证）：**

```bash
PUT /api/v1/users/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "newusername",
  "email": "newemail@example.com"
}
```

**删除用户（需要认证和管理员权限）：**

```bash
DELETE /api/v1/users/:id
Authorization: Bearer <token>
```

## 开发指南

### 添加新的模块

1. 在 `internal/model` 中定义数据模型
2. 在 `internal/dto` 中定义 DTO
3. 在 `internal/repository` 中实现数据访问层
4. 在 `internal/service` 中实现业务逻辑
5. 在 `internal/api/handler` 中实现 HTTP 处理器
6. 在 `internal/api/router` 中注册路由

### 运行测试

```bash
make test
```

生成测试覆盖率报告：

```bash
make test-coverage
```

### 代码检查

```bash
make lint
```

### 格式化代码

```bash
make fmt
```

### 构建应用

```bash
make build
```

编译后的二进制文件将在 `bin/server`。

## 配置说明

配置文件位于 `config/config.yaml`，支持以下配置项：

- **server**: 服务器配置（端口、模式、超时等）
- **database**: 数据库配置
- **redis**: Redis 配置
- **jwt**: JWT 配置（密钥、过期时间）

生产环境建议使用环境变量覆盖敏感配置：

```bash
export DB_PASSWORD=your_db_password
export JWT_SECRET=your_jwt_secret
```

## 中间件

项目包含以下中间件：

### 基础中间件

- **Logger**: 请求日志记录
- **Recovery**: Panic 恢复
- **CORS**: 跨域资源共享
- **Auth**: JWT 认证
- **RateLimit**: 限流
- **SecurityHeaders**: 安全响应头

### 高级中间件

- **Validate**: 通用 JSON 验证中间件
- **Pprof**: 性能分析工具（可配置）
- **Sentry**: 错误监控（可配置）
- **Tracing**: OpenTelemetry 分布式追踪（可配置）

> 详细配置和使用方法请参考 [高级功能指南](docs/FEATURES.md)

## Docker 部署

### 构建镜像

```bash
make docker-build
```

### 运行容器

```bash
make docker-run
```

### 使用 Docker Compose

```bash
docker-compose up -d
```

这将启动以下服务：

- 应用服务（端口 8080）
- PostgreSQL 数据库（端口 5432）
- Redis 缓存（端口 6379）

## 最佳实践

本项目遵循以下最佳实践：

1. **分层架构** - 清晰的职责分离，便于测试和维护
2. **依赖注入** - 使用构造函数注入，提高可测试性
3. **接口抽象** - Service 和 Repository 层使用接口定义
4. **错误处理** - 统一的错误响应格式
5. **参数验证** - 使用 validator 进行参数验证
6. **安全实践** - 密码加密、JWT 认证、安全响应头
7. **日志记录** - 结构化日志，便于问题排查
8. **优雅关闭** - 处理完现有请求后再关闭服务
9. **配置管理** - 使用配置文件和环境变量
10. **容器化** - 提供 Docker 支持，便于部署

## 技术栈

- **框架**: Gin
- **ORM**: GORM
- **日志**: Zap
- **配置**: Viper
- **JWT**: golang-jwt
- **验证**: validator
- **数据库**: PostgreSQL
- **缓存**: Redis（可选）

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

## 参考资料

- [Gin 官方文档](https://gin-gonic.com/)
- [GORM 官方文档](https://gorm.io/)
