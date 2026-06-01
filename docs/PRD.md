# MinimalTreeHole - 产品需求文档 (PRD)

**版本**: v1.0.0  
**创建时间**: 2026-06-02  
**状态**: 初始化

## 项目概述

**项目名称**: MinimalTreeHole (极简树洞测试应用)

**项目描述**: 这是一个用于验证全栈基础 CRUD 链路、并发读写以及基础安全策略的最小可行性产品（MVP）。应用允许用户无需注册即可发布匿名留言，并能查看他人的留言及进行简单的点赞互动。该项目作为架构探针，主要用于跑通代码编写到目标服务器自动化部署的全流程。

## 功能需求

### 核心功能

#### F1: 匿名发布机制
- 用户在前端提交纯文本留言（限制 500 字符以内）
- 后端接收数据，自动生成时间戳
- 数据存入 PostgreSQL 数据库
- 验收标准：
  - 前端输入框限制 500 字符
  - 提交成功后清空输入框并刷新列表
  - 后端返回 201 Created 状态码

#### F2: 公共信息流
- 主页以时间倒序（最新在前）展示所有树洞留言
- 后端接口支持基础的分页（Pagination）或基于游标的加载（Cursor-based loading）
- 验收标准：
  - 首次加载显示最新 20 条留言
  - 支持滚动加载更多
  - 每条留言显示内容、时间戳、抱抱计数

#### F3: 轻量级互动
- 每条留言附带一个"抱抱(+1)"按钮
- 用户点击后，该条留言的计数器实时 +1
- 验证数据库的并发更新操作及前端状态同步
- 验收标准：
  - 点击后立即更新前端显示
  - 后端正确处理并发点击（无计数丢失）
  - 同一 IP 可重复点击（无限制）

## 非功能需求

### 性能要求
- 核心信息流接口（Feed API）响应时间需小于 200ms
- 系统需支持基础的 100 并发请求
- 数据库查询需添加索引优化

### 安全要求
- **XSS 防御**: 对用户提交的文本进行转义清理
- **API 频率限制**: 基于 IP 限制每分钟最多发布 3 条留言，防止机器人恶意刷库
- **输入验证**: 后端验证留言长度、内容格式

### 可用性要求
- 依赖 Docker Compose 进行编排
- 配置异常自动重启策略（restart: always）
- 保证单节点高可用

## 技术栈

### 前端
- **框架**: React 18+
- **构建工具**: Vite
- **样式**: TailwindCSS
- **HTTP 客户端**: Axios 或 Fetch API

### 后端
- **语言**: Go 1.21+
- **框架**: Gin 或 Fiber
- **数据库驱动**: pgx (PostgreSQL)
- **限流**: 内存限流或 Redis

### 数据库
- **主数据库**: PostgreSQL 15+
- **缓存/限流**: Redis 7+ (可选，可先用 Go 内存限流)

### 基础设施
- **容器化**: Docker
- **编排**: Docker Compose
- **反向代理**: Nginx (可选)

## 环境配置

### DEV-HOST
- **位置**: 当前 Claude Code 环境 (/root)
- **权限**: AI 可直接操作
- **用途**: 代码开发、测试、Git 管理、编写 Dockerfile 与部署脚本

### TARGET-HOST
- **位置**: 43.106.137.193
- **SSH 信息**:
  - Host: 43.106.137.193
  - Port: 22
  - User: root
  - Password: 在 "/root/share_workspace/doc/test server pw.txt" 里
- **部署方式**: Deploy Agent 直接 SSH 操作（或通过 SSH 将 docker-compose 文件与构建产物推送到目标机并执行 `docker-compose up -d`）

## 数据模型

### Message (留言表)
```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL CHECK (char_length(content) <= 500),
    hug_count INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ip_address INET
);

CREATE INDEX idx_messages_created_at ON messages(created_at DESC);
```

## API 设计

### POST /api/messages
创建新留言

**请求体**:
```json
{
  "content": "string (max 500 chars)"
}
```

**响应** (201 Created):
```json
{
  "id": 1,
  "content": "string",
  "hug_count": 0,
  "created_at": "2026-06-02T10:00:00Z"
}
```

### GET /api/messages
获取留言列表

**查询参数**:
- `limit`: 每页数量 (默认 20)
- `cursor`: 游标 (上一页最后一条的 ID)

**响应** (200 OK):
```json
{
  "messages": [
    {
      "id": 1,
      "content": "string",
      "hug_count": 5,
      "created_at": "2026-06-02T10:00:00Z"
    }
  ],
  "next_cursor": 123
}
```

### POST /api/messages/:id/hug
给留言点赞

**响应** (200 OK):
```json
{
  "id": 1,
  "hug_count": 6
}
```

## 验收标准

### 功能验收
- [ ] 用户可以发布留言（500 字符以内）
- [ ] 主页显示所有留言（时间倒序）
- [ ] 支持分页或游标加载
- [ ] 点击"抱抱"按钮计数器 +1
- [ ] 前端实时更新显示

### 性能验收
- [ ] Feed API 响应时间 < 200ms
- [ ] 支持 100 并发请求无错误

### 安全验收
- [ ] XSS 防御生效（提交 `<script>alert(1)</script>` 被转义）
- [ ] 频率限制生效（1 分钟内第 4 条留言被拒绝）

### 部署验收
- [ ] Docker Compose 一键启动
- [ ] 服务异常自动重启
- [ ] 前端可通过浏览器访问
- [ ] 后端 API 可正常调用

## 里程碑

- **M1**: 后端 API 开发完成（CRUD + 限流）
- **M2**: 前端页面开发完成（发布 + 列表 + 点赞）
- **M3**: Docker Compose 编排完成
- **M4**: DEV-HOST 本地测试通过
- **M5**: TARGET-HOST 部署成功
- **M6**: UAT 验收通过

## 风险与假设

### 风险
- TARGET-HOST 环境未知（Docker/Git 是否已安装）
- 并发点赞可能导致计数不准确
- 无用户认证可能被恶意刷库

### 假设
- TARGET-HOST 有公网 IP 和开放端口
- PostgreSQL 和 Redis 可通过 Docker 部署
- 前端静态文件可通过 Nginx 或后端直接服务

## 后续优化方向

- 添加用户认证（可选）
- 引入 WebSocket 实时推送新留言
- 添加留言举报和审核机制
- 优化前端性能（虚拟滚动）
- 添加监控和日志系统
