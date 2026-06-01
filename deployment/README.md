# MinimalTreeHole 部署指南

## 快速启动

### 1. 配置环境变量
```bash
cd deployment
cp .env.example .env
# 编辑 .env 文件，修改数据库密码等配置
```

### 2. 启动所有服务
```bash
docker-compose up -d
```

### 3. 查看日志
```bash
docker-compose logs -f
```

### 4. 访问应用
- 前端: http://localhost
- 后端API: http://localhost:8080/api
- 健康检查: http://localhost:8080/health

## 服务管理

### 停止服务
```bash
docker-compose down
```

### 重启服务
```bash
docker-compose restart
```

### 查看服务状态
```bash
docker-compose ps
```

### 清理数据（谨慎操作）
```bash
docker-compose down -v
```

## 数据库迁移

数据库迁移脚本位于 `migrations/` 目录，会在 PostgreSQL 容器首次启动时自动执行。

如需手动执行：
```bash
docker exec -i treehole-db psql -U treehole -d treehole < migrations/001_create_messages.sql
docker exec -i treehole-db psql -U treehole -d treehole < migrations/002_create_indexes.sql
```

## 生产环境部署建议

1. **修改默认密码**: 在 `.env` 中设置强密码
2. **配置域名**: 修改 Nginx 配置中的 `server_name`
3. **启用 HTTPS**: 使用 Let's Encrypt 或其他 SSL 证书
4. **调整限流**: 根据实际需求修改 `RATE_LIMIT`
5. **备份数据**: 定期备份 PostgreSQL 数据卷

## 故障排查

### 查看后端日志
```bash
docker-compose logs backend
```

### 查看数据库日志
```bash
docker-compose logs postgres
```

### 进入容器调试
```bash
docker exec -it treehole-backend sh
docker exec -it treehole-db psql -U treehole -d treehole
```
