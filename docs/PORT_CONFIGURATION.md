# MinimalTreeHole 端口配置

## 当前端口分配

| 服务 | DEV-HOST | TARGET-HOST | 说明 |
|---|---|---|---|
| 前端 | 8001 | 8001 | 已确认可用 |
| 后端 | 8080 | 8082 | 8080 被 prerp-api 占用，改用 8082 |
| 数据库 | 5434 | 5434 | 已确认可用 |

## 访问地址

### DEV-HOST（本地测试）
- 前端：http://localhost:8001
- 后端 API：http://localhost:8080/api/health

### TARGET-HOST（生产部署）
- 前端：http://43.106.137.193:8001
- 后端 API：http://43.106.137.193:8082/api/health

## 端口冲突历史

### 2026-06-02
- **问题**：TARGET-HOST 的 8080 端口被 prerp-api (PID 71222) 占用
- **解决方案**：修改后端端口为 8082
- **影响**：前端需要调整 API 请求地址（Nginx 反向代理已配置）

## 配置文件

### docker-compose.yml (TARGET-HOST)
```yaml
backend:
  ports:
    - "8082:8080"  # 外部 8082 映射到容器内 8080
```

### nginx.conf (前端)
```nginx
location /api {
    proxy_pass http://backend:8080;  # 容器内部通信仍使用 8080
}
```

## 回滚方案

如果需要恢复 8080 端口：
1. 停止或迁移 prerp-api 服务
2. 修改 docker-compose.yml 恢复 `8080:8080`
3. 重启容器
