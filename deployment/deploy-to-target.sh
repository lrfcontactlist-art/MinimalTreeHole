#!/bin/bash
set -e

echo "=========================================="
echo "MinimalTreeHole 部署脚本"
echo "Commit: 6aaacb2"
echo "=========================================="
echo ""

# 1. 创建项目目录
echo "## 1. 创建项目目录"
mkdir -p /opt/MinimalTreeHole/{backend,frontend,deployment/migrations}
cd /opt/MinimalTreeHole
echo "✅ 项目目录创建完成"
echo ""

# 2. 检查 Docker
echo "## 2. 检查 Docker"
if ! command -v docker &> /dev/null; then
    echo "❌ Docker 未安装，正在安装..."
    curl -fsSL https://get.docker.com | sh
    systemctl start docker
    systemctl enable docker
    echo "✅ Docker 安装完成"
else
    echo "✅ Docker 已安装: $(docker --version)"
fi
echo ""

# 3. 检查 Docker Compose
echo "## 3. 检查 Docker Compose"
if docker compose version &> /dev/null; then
    echo "✅ Docker Compose 已安装: $(docker compose version)"
elif command -v docker-compose &> /dev/null; then
    echo "✅ docker-compose 已安装: $(docker-compose --version)"
else
    echo "❌ Docker Compose 未安装，正在安装..."
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
    echo "✅ Docker Compose 安装完成"
fi
echo ""

# 4. 停止旧容器（如果存在）
echo "## 4. 停止旧容器"
cd /opt/MinimalTreeHole/deployment
docker compose down 2>/dev/null || docker-compose down 2>/dev/null || echo "无旧容器需要停止"
echo ""

# 5. 启动服务
echo "## 5. 启动服务"
docker compose up -d --build || docker-compose up -d --build
echo ""

# 6. 等待服务启动
echo "## 6. 等待服务启动（30秒）"
sleep 30
echo ""

# 7. 检查容器状态
echo "## 7. 检查容器状态"
docker compose ps || docker-compose ps
echo ""

# 8. 健康检查
echo "## 8. 健康检查"
echo "后端健康检查:"
curl -s http://localhost:8080/api/health || echo "❌ 后端健康检查失败"
echo ""
echo "前端访问检查:"
curl -s http://localhost:8001 | head -5 || echo "❌ 前端访问失败"
echo ""

# 9. 查看日志
echo "## 9. 查看最近日志"
docker compose logs --tail=20 || docker-compose logs --tail=20
echo ""

echo "=========================================="
echo "部署完成！"
echo "前端访问: http://43.106.137.193:8001"
echo "后端 API: http://43.106.137.193:8080"
echo "=========================================="
