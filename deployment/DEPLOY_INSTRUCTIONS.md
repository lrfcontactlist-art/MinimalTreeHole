# MinimalTreeHole 部署指引

## 当前进度：S08.0 ENV-PROBE（环境探测）

### 步骤 1：登录 TARGET-HOST

```bash
ssh root@43.106.137.193
```

密码见：`/root/share_workspace/doc/test server pw.txt`

### 步骤 2：执行环境探测脚本

在 TARGET-HOST 上直接复制粘贴以下完整脚本并执行：

```bash
bash << 'SCRIPTEOF'
echo "=========================================="
echo "TARGET-HOST 环境探测脚本"
echo "=========================================="
echo ""
echo "## 1. 系统信息"
echo "操作系统: $(uname -s)"
echo "内核版本: $(uname -r)"
echo "架构: $(uname -m)"
echo "主机名: $(hostname)"
echo ""
echo "## 2. 磁盘空间"
df -h | grep -E "Filesystem|/$|/home|/var"
echo ""
echo "## 3. 内存信息"
free -h
echo ""
echo "## 4. Docker 检查"
if command -v docker &> /dev/null; then
    echo "✅ Docker 已安装"
    docker --version
    echo "Docker 服务状态:"
    systemctl is-active docker 2>/dev/null || service docker status 2>/dev/null | head -3
else
    echo "❌ Docker 未安装"
fi
echo ""
echo "## 5. Docker Compose 检查"
if command -v docker-compose &> /dev/null; then
    echo "✅ docker-compose 已安装"
    docker-compose --version
elif docker compose version &> /dev/null; then
    echo "✅ docker compose (plugin) 已安装"
    docker compose version
else
    echo "❌ Docker Compose 未安装"
fi
echo ""
echo "## 6. Git 检查"
if command -v git &> /dev/null; then
    echo "✅ Git 已安装"
    git --version
else
    echo "❌ Git 未安装"
fi
echo ""
echo "## 7. 端口占用检查"
echo "检查关键端口..."
for port in 80 443 3000 5432 8080 8001; do
    if netstat -tuln 2>/dev/null | grep -q ":$port " || ss -tuln 2>/dev/null | grep -q ":$port "; then
        echo "⚠️  端口 $port 已被占用"
    else
        echo "✅ 端口 $port 可用"
    fi
done
echo ""
echo "## 8. 网络连接检查"
if ping -c 1 8.8.8.8 &> /dev/null; then
    echo "✅ 外网连接正常"
else
    echo "❌ 外网连接失败"
fi
echo ""
echo "## 9. 当前运行的容器"
if command -v docker &> /dev/null; then
    docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" 2>/dev/null || echo "无运行中的容器"
else
    echo "Docker 未安装，跳过"
fi
echo ""
echo "=========================================="
echo "探测完成！请将以上输出复制回 DEV-HOST"
echo "=========================================="
SCRIPTEOF
```

### 步骤 3：复制输出

将 TARGET-HOST 上的完整输出复制，然后粘贴回 DEV-HOST 的 Claude Code 会话中。

---

## 下一步

收到环境探测结果后，将自动进入 S08.1 DEPLOY-SCRIPT 阶段，生成针对性的部署脚本。
