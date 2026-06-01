# MinimalTreeHole

极简树洞测试应用 - 全栈 CRUD 链路验证项目

## 项目简介

MinimalTreeHole 是一个用于验证全栈基础 CRUD 链路、并发读写以及基础安全策略的最小可行性产品（MVP）。应用允许用户无需注册即可发布匿名留言，并能查看他人的留言及进行简单的点赞互动。

## 核心功能

- **匿名发布**: 用户可提交纯文本留言（限制 500 字符以内）
- **公共信息流**: 主页以时间倒序展示所有树洞留言，支持分页加载
- **轻量级互动**: 每条留言附带"抱抱(+1)"按钮，点击后计数器实时 +1

## 技术栈

- **前端**: React 18 + Vite + TailwindCSS
- **后端**: Go 1.21+ + Gin/Fiber
- **数据库**: PostgreSQL 15+
- **基础设施**: Docker + Docker Compose

## 项目结构

```
MinimalTreeHole/
├── frontend/          # React 前端应用
├── backend/           # Go 后端服务
├── deployment/        # Docker Compose 和部署脚本
└── docs/              # 项目文档
    ├── PRD.md         # 产品需求文档
    ├── AOCI.md        # 架构组件索引
    ├── CHANGELOG.md   # 变更日志
    └── DEPLOYMENT.md  # 部署文档
```

## 快速开始

### 开发环境

```bash
# 克隆仓库
git clone <repository-url>
cd MinimalTreeHole

# 启动开发环境
docker-compose -f deployment/docker-compose.yml up -d
```

### 访问应用

- 前端: http://localhost:3000
- 后端 API: http://localhost:8080/api

## 开发方法论

本项目使用 **RDM-M-Swarm-DS-v2** 方法论进行开发：

- **双服务器隔离**: DEV-HOST（AI 可操作）+ TARGET-HOST（人工搬运部署）
- **AOCI 索引驱动**: 先建立系统地图，再精确定位影响范围
- **多 Agent 协同**: 主控决策 + 专职 Agent 执行（Code/Test/AOCI/Deploy）
- **结构化验证**: 机械验证代替信任度，确保改动可靠

## 文档

- [产品需求文档 (PRD)](docs/PRD.md)
- [架构组件索引 (AOCI)](docs/AOCI.md)
- [部署文档 (DEPLOYMENT)](docs/DEPLOYMENT.md)
- [变更日志 (CHANGELOG)](docs/CHANGELOG.md)

## 安全特性

- **XSS 防御**: 对用户提交的文本进行转义清理
- **API 频率限制**: 基于 IP 限制每分钟最多发布 3 条留言
- **输入验证**: 后端验证留言长度、内容格式

## 性能指标

- Feed API 响应时间 < 200ms
- 支持 100 并发请求
- 数据库查询优化（索引）

## 许可证

MIT License

## 联系方式

GitHub: lrfcontactlist-art
