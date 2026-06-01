#!/bin/bash
set -e

echo "=========================================="
echo "MinimalTreeHole GitHub 部署脚本"
echo "Commit: 6aaacb2"
echo "仓库: https://github.com/lrfcontactlist-art/MinimalTreeHole"
echo "=========================================="
echo ""

REPO_URL="https://github.com/lrfcontactlist-art/MinimalTreeHole.git"
DEPLOY_DIR="/opt/MinimalTreeHole"
COMMIT_ID="6aaacb2"

# 1. 检查 Git
echo "## 1. 检查 Git"
if ! command -v git &> /dev/null; then
    echo "安装 Git..."
    apt-get update && apt-get install -y git || yum install -y git
fi
echo "✅ Git: $(git --version)"
echo ""

# 2. 检查 Docker
echo "## 2. 检查 Docker"
if ! command -v docker &> /dev/null; then
    echo "安装 Docker..."
    curl -fsSL https://get.docker.com | sh
    systemctl start docker
    systemctl enable docker
fi
echo "✅ Docker: $(docker --version)"
echo ""

# 3. 检查 Docker Compose
echo "## 3. 检查 Docker Compose"
if ! docker compose version &> /dev/null && ! command -v docker-compose &> /dev/null; then
    echo "安装 Docker Compose..."
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
fi
echo "✅ Docker Compose 已安装"
echo ""

# 4. 克隆或更新代码
echo "## 4. 克隆/更新代码"
if [ -d "$DEPLOY_DIR/.git" ]; then
    echo "更新现有仓库..."
    cd "$DEPLOY_DIR"
    git fetch origin
    git checkout "$COMMIT_ID"
else
    echo "克隆仓库..."
    rm -rf "$DEPLOY_DIR"
    git clone "$REPO_URL" "$DEPLOY_DIR"
    cd "$DEPLOY_DIR"
    git checkout "$COMMIT_ID"
fi
echo "✅ 代码已更新到 commit: $(git rev-parse --short HEAD)"
echo ""

# 5. 停止旧容器
echo "## 5. 停止旧容器"
cd "$DEPLOY_DIR/deployment"
docker compose down 2>/dev/null || docker-compose down 2>/dev/null || echo "无旧容器"
echo ""

# 6. 启动服务
echo "## 6. 启动服务"
docker compose up -d --build || docker-compose up -d --build
echo ""

# 7. 等待服务启动
echo "## 7. 等待服务启动（30秒）"
sleep 30
echo ""

# 8. 检查容器状态
echo "## 8. 检查容器状态"
docker compose ps || docker-compose ps
echo ""

# 9. 健康检查
echo "## 9. 健康检查"
echo "后端健康检查:"
curl -s http://localhost:8080/api/health || echo "❌ 后端健康检查失败"
echo ""
echo "前端访问检查:"
curl -s http://localhost:8001 | head -5 || echo "❌ 前端访问失败"
echo ""

# 10. 查看日志
echo "## 10. 最近日志"
docker compose logs --tail=20 || docker-compose logs --tail=20
echo ""

echo "=========================================="
echo "✅ 部署完成！"
echo ""
echo "访问地址:"
echo "  前端: http://$(curl -s ifconfig.me):8001"
echo "  后端: http://$(curl -s ifconfig.me):8080"
echo ""
echo "本地测试:"
echo "  curl http://localhost:8080/api/health"
echo "  curl http://localhost:8001"
echo "=========================================="
