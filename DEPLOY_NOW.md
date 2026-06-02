# 🚀 MinimalTreeHole 立即部署指南

## 📦 部署包已就绪

**文件位置**: `/tmp/MinimalTreeHole-deploy.tar.gz` (17MB)  
**Commit ID**: 6aaacb2

---

## 🎯 快速部署（3 步完成）

### 步骤 1：传输部署包到 TARGET-HOST

在 **DEV-HOST** 上执行：

```bash
scp /tmp/MinimalTreeHole-deploy.tar.gz root@43.106.137.193:/tmp/
```

输入密码：`[JAar;H2x6Tq-P_]`

---

### 步骤 2：登录 TARGET-HOST

```bash
ssh root@43.106.137.193
```

输入密码：`[JAar;H2x6Tq-P_]`

---

### 步骤 3：解压并部署

在 **TARGET-HOST** 上执行：

```bash
# 解压到 /opt
cd /opt
tar -xzf /tmp/MinimalTreeHole-deploy.tar.gz

# 执行部署脚本
cd MinimalTreeHole/deployment
bash deploy-to-target.sh
```

部署脚本会自动：
- ✅ 检查并安装 Docker/Docker Compose
- ✅ 停止旧容器
- ✅ 构建并启动新容器
- ✅ 执行健康检查

---

## 🔍 验证部署

### 在 TARGET-HOST 上测试

```bash
# 检查容器状态
docker compose ps

# 测试后端 API
curl http://localhost:8080/api/health

# 创建测试留言
curl -X POST http://localhost:8080/api/messages \
  -H "Content-Type: application/json" \
  -d '{"content":"Hello TreeHole!"}'

# 获取留言列表
curl http://localhost:8080/api/messages

# 测试前端
curl http://localhost:8001 | head -10
```

### 在浏览器中访问

- **前端**: http://43.106.137.193:8001
- **后端 API**: http://43.106.137.193:8080/api/health

---

## 📊 UAT 测试清单

在浏览器中完成以下测试：

- [ ] 打开前端页面，看到"树洞"标题
- [ ] 输入留言（500 字符以内），点击发布
- [ ] 看到留言出现在列表顶部
- [ ] 点击"抱抱"按钮，计数器 +1
- [ ] 快速发布 4 条留言，第 4 条被拒绝（频率限制）
- [ ] 测试 XSS 防御：输入 `<script>alert(1)</script>`，显示为转义文本

---

## 🛠️ 故障排查

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

### 检查端口占用
```bash
netstat -tuln | grep -E "8080|8001|5434"
```

---

## 📝 服务信息

| 服务 | 容器名 | 端口映射 | 状态检查 |
|---|---|---|---|
| 前端 | treehole-frontend | 8001:80 | `curl http://localhost:8001` |
| 后端 | treehole-backend | 8080:8080 | `curl http://localhost:8080/api/health` |
| 数据库 | treehole-db | 5434:5432 | `docker exec treehole-db psql -U treehole -c "SELECT 1"` |

---

## ✅ 完成后

部署成功并通过 UAT 测试后，在 Claude Code 会话中回复：

```
部署成功！UAT 测试通过。
```

我将更新任务状态并生成最终报告。

---

**开始部署吧！** 🚀
