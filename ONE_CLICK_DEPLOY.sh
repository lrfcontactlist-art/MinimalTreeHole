#!/bin/bash
set -e

echo "=========================================="
echo "MinimalTreeHole 一键部署脚本"
echo "=========================================="
echo ""

TARGET_HOST="43.106.137.193"
TARGET_USER="root"
TARGET_PASS='[JAar;H2x6Tq-P_]'
DEPLOY_PKG="/tmp/MinimalTreeHole-deploy.tar.gz"

# 检查 sshpass
if ! command -v sshpass &> /dev/null; then
    echo "❌ sshpass 未安装，正在安装..."
    apt-get update && apt-get install -y sshpass || yum install -y sshpass
fi

echo "## 步骤 1: 传输部署包到 TARGET-HOST"
sshpass -p "$TARGET_PASS" scp -o StrictHostKeyChecking=no "$DEPLOY_PKG" "$TARGET_USER@$TARGET_HOST:/tmp/"
echo "✅ 部署包传输完成"
echo ""

echo "## 步骤 2: 在 TARGET-HOST 上解压并部署"
sshpass -p "$TARGET_PASS" ssh -o StrictHostKeyChecking=no "$TARGET_USER@$TARGET_HOST" << 'REMOTE_EOF'
set -e

echo "解压部署包..."
cd /opt
rm -rf MinimalTreeHole-backup-* 2>/dev/null || true
mv MinimalTreeHole MinimalTreeHole-backup-$(date +%Y%m%d-%H%M%S) 2>/dev/null || true
tar -xzf /tmp/MinimalTreeHole-deploy.tar.gz

echo "检查 Docker..."
if ! command -v docker &> /dev/null; then
    echo "安装 Docker..."
    curl -fsSL https://get.docker.com | sh
    systemctl start docker
    systemctl enable docker
fi

echo "检查 Docker Compose..."
if ! docker compose version &> /dev/null && ! command -v docker-compose &> /dev/null; then
    echo "安装 Docker Compose..."
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
fi

echo "启动服务..."
cd /opt/MinimalTreeHole/deployment
docker compose down 2>/dev/null || docker-compose down 2>/dev/null || true
docker compose up -d --build || docker-compose up -d --build

echo "等待服务启动（30秒）..."
sleep 30

echo "检查容器状态..."
docker compose ps || docker-compose ps

echo ""
echo "=========================================="
echo "部署完成！"
echo "=========================================="
REMOTE_EOF

echo ""
echo "## 步骤 3: 健康检查"
echo "后端健康检查:"
sshpass -p "$TARGET_PASS" ssh -o StrictHostKeyChecking=no "$TARGET_USER@$TARGET_HOST" "curl -s http://localhost:8080/api/health"
echo ""

echo "前端访问检查:"
sshpass -p "$TARGET_PASS" ssh -o StrictHostKeyChecking=no "$TARGET_USER@$TARGET_HOST" "curl -s http://localhost:8001 | head -5"
echo ""

echo "## 步骤 4: 功能测试"
echo "创建测试留言:"
sshpass -p "$TARGET_PASS" ssh -o StrictHostKeyChecking=no "$TARGET_USER@$TARGET_HOST" \
  "curl -s -X POST http://localhost:8080/api/messages -H 'Content-Type: application/json' -d '{\"content\":\"测试留言\"}'"
echo ""

echo "获取留言列表:"
sshpass -p "$TARGET_PASS" ssh -o StrictHostKeyChecking=no "$TARGET_USER@$TARGET_HOST" \
  "curl -s http://localhost:8080/api/messages"
echo ""

echo "=========================================="
echo "✅ 部署和测试完成！"
echo ""
echo "访问地址:"
echo "  前端: http://43.106.137.193:8001"
echo "  后端: http://43.106.137.193:8080"
echo ""
echo "请在浏览器中进行 UAT 测试！"
echo "=========================================="
