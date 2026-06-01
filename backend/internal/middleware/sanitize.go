package middleware

import (
	"bytes"
	"encoding/json"
	"html"
	"io"

	"github.com/gin-gonic/gin"
)

func SanitizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 只处理 POST /api/messages
		if c.Request.Method == "POST" && c.Request.URL.Path == "/api/messages" {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.Next()
				return
			}
			
			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				// 恢复 body 供后续处理
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				c.Next()
				return
			}
			
			// 清理 content 字段
			if content, ok := data["content"].(string); ok {
				data["content"] = html.EscapeString(content)
			}
			
			// 重新序列化
			sanitized, err := json.Marshal(data)
			if err != nil {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				c.Next()
				return
			}
			
			// 替换 request body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(sanitized))
			c.Request.ContentLength = int64(len(sanitized))
		}
		
		c.Next()
	}
}
