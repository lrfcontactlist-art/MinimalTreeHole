# MinimalTreeHole 手动部署指南

## 部署方式：手动传输文件

由于安全限制，需要手动将文件传输到 TARGET-HOST。

---

## 步骤 1：打包项目文件

在 DEV-HOST 上执行：

```bash
cd /root/MinimalTreeHole
tar -czf /tmp/MinimalTreeHole-deploy.tar.gz \
  backend/ \
  frontend/ \
  deployment/ \
  --exclude=backend/node_modules \
  --exclude=frontend/node_modules \
  --exclude=frontend/dist
```

---

## 步骤 2：传输到 TARGET-HOST

### 方式 A：使用 scp（推荐）

```bash
scp /tmp/MinimalTreeHole-deploy.tar.gz root@43.106.137.193:/tmp/
```

### 方式 B：使用其他工具
- SFTP 客户端（FileZilla, WinSCP）
- rsync
- 云存储中转

---

## 步骤 3：在 TARGET-HOST 上解压并部署

登录 TARGET-HOST：
```bash
ssh root@43.106.137.193
```

密码：`[JAar;H2x6Tq-P_]`

执行部署：
```bash
# 解压文件
cd /opt
tar -xzf /tmp/MinimalTreeHole-deploy.tar.gz
mv MinimalTreeHole MinimalTreeHole-backup-$(date +%Y%m%d) 2>/dev/null || true
tar -xzf /tmp/MinimalTreeHole-deploy.tar.gz

# 进入部署目录
cd /opt/MinimalTreeHole/deployment

# 执行部署脚本
bash deploy-to-target.sh
```

---

## 步骤 4：验证部署

### 检查容器状态
```bash
docker compose ps
```

### 测试后端 API
```bash
# 健康检查
curl http://localhost:8080/api/health

# 创建留言
curl -X POST http://localhost:8080/api/messages \
  -H "Content-Type: application/json" \
  -d '{"content":"测试留言"}'

# 获取留言列表
curl http://localhost:8080/api/messages
```

### 测试前端
```bash
curl http://localhost:8001
```

### 浏览器访问
- 前端：http://43.106.137.193:8001
- 后端 API：http://43.106.137.193:8080/api/health

---

## 故障排查

### 查看日志
```bash
cd /opt/MinimalTreeHole/deployment
docker compose logs -f
```

### 重启服务
```bash
docker compose restart
```

### 完全重新部署
```bash
docker compose down
docker compose up -d --build
```

---

## 端口说明

- **前端**: 8001 → 80（容器内）
- **后端**: 8080 → 8080（容器内）
- **数据库**: 5434 → 5432（容器内）

如果端口被占用，修改 `deployment/docker-compose.yml` 中的端口映射。
