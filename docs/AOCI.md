# MinimalTreeHole - AOCI 索引

**版本**: v1.0.0  
**创建时间**: 2026-06-02  
**状态**: 待构建（S02.1 AOCI-BUILD）

## 索引说明

AOCI (Architecture-Oriented Component Index) 是系统组件的结构化索引，用于快速定位影响范围。

每个条目包含：
- **F** (File): 文件路径
- **R** (Role): 组件职责
- **A** (API): 对外接口
- **S** (Stack): 技术栈细节（框架、返回类型、函数签名）

## 前端组件

### FE-001: 主应用入口
- **F**: `frontend/src/main.tsx`
- **R**: React 应用挂载点，路由配置
- **A**: 无（入口文件）
- **S**: React 18 + ReactDOM.createRoot() + Vite

### FE-002: 留言列表组件
- **F**: `frontend/src/components/MessageList.tsx`
- **R**: 展示留言列表，支持分页加载
- **A**: Props: `messages: Message[]`, `onLoadMore: () => void`
- **S**: React.FC<Props>, TailwindCSS, 调用 GET /api/messages

### FE-003: 留言发布组件
- **F**: `frontend/src/components/MessageForm.tsx`
- **R**: 留言输入框，字符计数，提交逻辑
- **A**: Props: `onSubmit: (content: string) => Promise<void>`
- **S**: React.FC<Props>, useState hook, 调用 POST /api/messages

### FE-004: 留言卡片组件
- **F**: `frontend/src/components/MessageCard.tsx`
- **R**: 单条留言展示，抱抱按钮
- **A**: Props: `message: Message`, `onHug: (id: number) => Promise<void>`
- **S**: React.FC<Props>, 调用 POST /api/messages/:id/hug

### FE-005: API 客户端
- **F**: `frontend/src/api/client.ts`
- **R**: 封装所有后端 API 调用
- **A**: `fetchMessages(cursor?: number): Promise<MessageResponse>`, `createMessage(content: string): Promise<Message>`, `hugMessage(id: number): Promise<Message>`
- **S**: Axios 或 Fetch API, 返回 Promise<T>, 基础 URL: /api

### FE-006: 类型定义
- **F**: `frontend/src/types/index.ts`
- **R**: TypeScript 类型定义
- **A**: `interface Message`, `interface MessageResponse`
- **S**: TypeScript interfaces

## 后端组件

### BE-001: 主程序入口
- **F**: `backend/cmd/main.go`
- **R**: 启动 HTTP 服务器，初始化数据库连接
- **A**: `func main()`
- **S**: Go 1.21+, Gin/Fiber 框架, 监听 :8080

### BE-002: 路由配置
- **F**: `backend/internal/router/router.go`
- **R**: 注册所有 API 路由和中间件
- **A**: `func SetupRouter(db *sql.DB) *gin.Engine`
- **S**: Gin Router, 返回 *gin.Engine, 挂载中间件链

### BE-003: 留言处理器
- **F**: `backend/internal/handler/message.go`
- **R**: 处理留言相关的 HTTP 请求
- **A**: `func CreateMessage(c *gin.Context)`, `func GetMessages(c *gin.Context)`, `func HugMessage(c *gin.Context)`
- **S**: Gin handlers (gin.HandlerFunc), 返回 JSON, 状态码 200/201/400/500

### BE-004: 留言服务层
- **F**: `backend/internal/service/message.go`
- **R**: 业务逻辑，调用数据库层
- **A**: `func CreateMessage(content, ip string) (*model.Message, error)`, `func ListMessages(limit int, cursor *int) ([]*model.Message, *int, error)`, `func IncrementHug(id int) (*model.Message, error)`
- **S**: Go service, 返回 (data, error), 包含业务验证

### BE-005: 留言数据层
- **F**: `backend/internal/repository/message.go`
- **R**: 数据库 CRUD 操作
- **A**: `func Insert(msg *model.Message) error`, `func FindAll(limit int, cursor *int) ([]*model.Message, error)`, `func UpdateHugCount(id int) error`
- **S**: pgx driver (*sql.DB), 返回 (data, error), 使用预编译语句

### BE-006: 数据库模型
- **F**: `backend/internal/model/message.go`
- **R**: 数据结构定义
- **A**: `type Message struct`
- **S**: Go struct with json tags, time.Time, sql.NullString

### BE-007: 限流中间件
- **F**: `backend/internal/middleware/ratelimit.go`
- **R**: IP 限流，防止恶意刷库
- **A**: `func RateLimitMiddleware() gin.HandlerFunc`
- **S**: Go middleware, sync.Map 存储 IP 计数, 每分钟 3 次限制

### BE-008: XSS 防御中间件
- **F**: `backend/internal/middleware/sanitize.go`
- **R**: 输入清理，防止 XSS 攻击
- **A**: `func SanitizeMiddleware() gin.HandlerFunc`
- **S**: Go middleware, html.EscapeString(), 处理 request body

### BE-009: 数据库连接
- **F**: `backend/internal/database/postgres.go`
- **R**: PostgreSQL 连接池管理
- **A**: `func NewPostgresDB(connStr string) (*sql.DB, error)`
- **S**: pgx driver, 返回 *sql.DB, 连接池配置 MaxOpenConns=25

### BE-010: 配置管理
- **F**: `backend/internal/config/config.go`
- **R**: 读取环境变量和配置文件
- **A**: `func LoadConfig() *Config`, `type Config struct`
- **S**: Go struct, os.Getenv(), 默认值处理

## 数据库组件

### DB-001: 留言表
- **F**: `deployment/migrations/001_create_messages.sql`
- **R**: 留言数据存储
- **A**: `CREATE TABLE messages (...)`
- **S**: PostgreSQL 15+, SERIAL PRIMARY KEY, TIMESTAMP WITH TIME ZONE

### DB-002: 索引
- **F**: `deployment/migrations/002_create_indexes.sql`
- **R**: 查询性能优化
- **A**: `CREATE INDEX idx_messages_created_at ON messages(created_at DESC)`
- **S**: PostgreSQL B-tree index, 支持倒序查询

## 部署组件

### DEPLOY-001: Docker Compose
- **F**: `deployment/docker-compose.yml`
- **R**: 多容器编排（前端、后端、数据库）
- **A**: services: frontend, backend, postgres
- **S**: Docker Compose v3.8+, networks, volumes, restart: always

### DEPLOY-002: 后端 Dockerfile
- **F**: `backend/Dockerfile`
- **R**: 后端镜像构建
- **A**: 多阶段构建（builder + runtime）
- **S**: golang:1.21-alpine, CGO_ENABLED=0, 二进制输出 /app/server

### DEPLOY-003: 前端 Dockerfile
- **F**: `frontend/Dockerfile`
- **R**: 前端镜像构建
- **A**: 多阶段构建（builder + nginx）
- **S**: node:20-alpine + nginx:alpine, 输出 /usr/share/nginx/html

### DEPLOY-004: Nginx 配置
- **F**: `deployment/nginx.conf`
- **R**: 前端静态文件服务，API 反向代理
- **A**: server 配置, location /api 代理到后端
- **S**: Nginx 1.24+, proxy_pass http://backend:8080

### DEPLOY-005: 环境变量模板
- **F**: `deployment/.env.example`
- **R**: 环境变量示例
- **A**: DATABASE_URL, BACKEND_PORT, FRONTEND_PORT
- **S**: KEY=VALUE 格式, 注释说明

### DEPLOY-006: 部署脚本
- **F**: `deployment/deploy.sh`
- **R**: 自动化部署到 TARGET-HOST
- **A**: SSH 连接, git pull, docker-compose up -d
- **S**: Bash script, 绑定 commit id, 错误处理

### DEPLOY-007: 健康检查脚本
- **F**: `deployment/healthcheck.sh`
- **R**: 验证服务状态
- **A**: HTTP 请求 /api/health, 数据库连接测试
- **S**: Bash script, curl + psql, 退出码 0/1

## 依赖关系

```
FE-002 (MessageList) → FE-004 (MessageCard) → FE-005 (API Client) → BE-003 (Handler)
FE-003 (MessageForm) → FE-005 (API Client) → BE-003 (Handler)
BE-003 (Handler) → BE-004 (Service) → BE-005 (Repository) → DB-001 (Table)
BE-002 (Router) → BE-007 (RateLimit) → BE-008 (Sanitize) → BE-003 (Handler)
DEPLOY-001 (Compose) → DEPLOY-002 (Backend Dockerfile) + DEPLOY-003 (Frontend Dockerfile)
```

## 更新日志

- 2026-06-02: 初始化 AOCI 索引结构（待 S02.1 AOCI-BUILD 完善技术栈细节）
