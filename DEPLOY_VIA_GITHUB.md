# 🚀 MinimalTreeHole GitHub 部署指南

## ✅ 代码已推送到 GitHub

**仓库地址**: https://github.com/lrfcontactlist-art/MinimalTreeHole  
**最新 Commit**: a229a83  
**部署方式**: TARGET-HOST 从 GitHub 克隆代码

---

## 📋 部署步骤（在 TARGET-HOST 执行）

### 步骤 1：登录 TARGET-HOST

```bash
ssh root@43.106.137.193
```

密码：`[JAar;H2x6Tq-P_]`

---

### 步骤 2：下载并执行部署脚本

在 TARGET-HOST 上执行：

```bash
# 下载部署脚本
curl -o /tmp/deploy.sh https://raw.githubusercontent.com/lrfcontactlist-art/MinimalTreeHole/master/deployment/deploy-from-github.sh

# 执行部署
bash /tmp/deploy.sh
```

**或者直接一行命令**：

```bash
curl -fsSL https://raw.githubusercontent.com/lrfcontactlist-art/MinimalTreeHole/master/deployment/deploy-from-github.sh | bash
```

---

## 🔍 部署脚本会自动完成

1. ✅ 检查并安装 Git
2. ✅ 检查并安装 Docker
3. ✅ 检查并安装 Docker Compose
4. ✅ 从 GitHub 克隆代码到 `/opt/MinimalTreeHole`
5. ✅ 停止旧容器
6. ✅ 构建并启动新容器
7. ✅ 执行健康检查
8. ✅ 输出访问地址

---

## 📊 验证部署

### 在 TARGET-HOST 上测试

```bash
# 检查容器状态
cd /opt/MinimalTreeHole/deployment
docker compose ps

# 测试后端 API
curl http://localhost:8080/api/health

# 创建测试留言
curl -X POST http://localhost:8080/api/messages \
  -H "Content-Type: application/json" \
  -d '{"content":"Hello from GitHub!"}'

# 获取留言列表
curl http://localhost:8080/api/messages

# 测试前端
curl http://localhost:8001 | head -10
```

### 在浏览器中访问

- **前端**: http://43.106.137.193:8001
- **后端 API**: http://43.106.137.193:8080/api/health

---

## 📝 UAT 测试清单

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

### 回滚到之前的版本
```bash
cd /opt/MinimalTreeHole
git checkout 6aaacb2  # 回滚到初始版本
cd deployment
docker compose down
docker compose up -d --build
```

---

## ✅ 完成后

部署成功并通过 UAT 测试后，在 Claude Code 会话中回复：

```
部署成功！UAT 测试通过。
```

---

**现在开始部署吧！** 🚀
