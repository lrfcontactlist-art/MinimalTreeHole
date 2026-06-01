# MinimalTreeHole - 部署文档

**版本**: v1.0.0  
**创建时间**: 2026-06-02  
**状态**: 待部署

## 环境架构

### DEV-HOST
- **位置**: /root/MinimalTreeHole
- **用途**: 代码开发、测试、Git 管理
- **权限**: AI 可直接操作

### TARGET-HOST
- **IP**: 43.106.137.193
- **SSH**: root@43.106.137.193:22
- **密码**: 见 /root/share_workspace/doc/test server pw.txt
- **用途**: 生产部署
- **权限**: 人工搬运部署

### GITHUB
- **仓库**: 待创建
- **用途**: 代码同步、版本管理、回滚支持

## 部署流程

### 阶段 1: DEV-HOST 开发
1. 代码开发完成
2. 本地测试通过
3. Git commit + push

### 阶段 2: 环境探测（S08.0）
1. 生成环境探测脚本
2. 人工在 TARGET-HOST 执行
3. 收集环境信息（Docker、Git、端口占用）

### 阶段 3: 部署脚本生成（S08.1）
1. 基于环境探测结果生成部署脚本
2. 脚本绑定 commit id
3. 包含备份、构建检查、健康检查

### 阶段 4: 人工部署（S09）
1. 将部署脚本复制到 TARGET-HOST
2. 执行部署脚本
3. 粘贴输出回 DEV-HOST

### 阶段 5: 部署审查（S10）
1. 审计部署输出
2. 确认服务状态
3. 检查健康检查结果

### 阶段 6: 自动验证（S11）
1. curl 验证关键接口
2. 检查返回状态码和响应体

### 阶段 7: 用户验收（S12）
1. 人类在浏览器手动测试
2. 验证 UI、交互、业务逻辑

## 端口规划

| 服务 | DEV-HOST | TARGET-HOST |
|---|---|---|
| 前端 | 3000 | 3000 |
| 后端 | 8080 | 8080 |
| PostgreSQL | 5432 | 5432 |
| Redis (可选) | 6379 | 6379 |

## 环境变量

### 后端环境变量
```bash
DATABASE_URL=postgres://user:pass@postgres:5432/treehole
BACKEND_PORT=8080
RATE_LIMIT_PER_MINUTE=3
```

### 前端环境变量
```bash
VITE_API_BASE_URL=/api
```

## Docker Compose 服务

### frontend
- 镜像: 基于 node:20-alpine + nginx:alpine
- 端口: 3000:80
- 依赖: backend

### backend
- 镜像: 基于 golang:1.21-alpine
- 端口: 8080:8080
- 依赖: postgres
- 环境变量: DATABASE_URL

### postgres
- 镜像: postgres:15-alpine
- 端口: 5432:5432
- 卷: postgres_data
- 环境变量: POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB

## 健康检查

### 后端健康检查
```bash
curl http://localhost:8080/api/health
# 期望: {"status":"ok"}
```

### 前端健康检查
```bash
curl http://localhost:3000
# 期望: HTTP 200, HTML 内容
```

### 数据库健康检查
```bash
docker exec -it postgres psql -U user -d treehole -c "SELECT 1"
# 期望: 返回 1
```

## 回滚方案

### 方案 1: Git 回滚
```bash
git checkout <previous-commit-id>
docker-compose down
docker-compose up -d --build
```

### 方案 2: Docker 镜像回滚
```bash
docker-compose down
docker tag treehole-backend:latest treehole-backend:backup
docker tag treehole-backend:previous treehole-backend:latest
docker-compose up -d
```

## 监控指标

- 后端 API 响应时间 < 200ms
- 数据库连接池使用率 < 80%
- 容器重启次数 = 0
- 磁盘使用率 < 80%

## 故障排查

### 后端无法启动
1. 检查数据库连接: `docker logs backend`
2. 检查环境变量: `docker exec backend env`
3. 检查端口占用: `netstat -tuln | grep 8080`

### 前端无法访问
1. 检查 Nginx 配置: `docker exec frontend nginx -t`
2. 检查后端代理: `curl http://backend:8080/api/health`
3. 检查容器状态: `docker ps`

### 数据库连接失败
1. 检查 PostgreSQL 状态: `docker logs postgres`
2. 检查网络连接: `docker network inspect treehole_default`
3. 检查数据库用户权限: `psql -U user -d treehole`

## 安全注意事项

- 不要将 .env 文件提交到 Git
- 定期更新 Docker 镜像
- 限制 PostgreSQL 端口仅内部访问
- 启用 HTTPS（生产环境）
- 定期备份数据库

## 备份策略

### 数据库备份
```bash
docker exec postgres pg_dump -U user treehole > backup_$(date +%Y%m%d_%H%M%S).sql
```

### 代码备份
```bash
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

## 更新日志

- 2026-06-02: 初始化部署文档
