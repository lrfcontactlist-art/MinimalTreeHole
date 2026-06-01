package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	rl := &rateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
	
	// 定期清理过期记录
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			rl.cleanup()
		}
	}()
	
	return rl
}

func (rl *rateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	for ip, times := range rl.requests {
		var valid []time.Time
		for _, t := range times {
			if now.Sub(t) < rl.window {
				valid = append(valid, t)
			}
		}
		if len(valid) == 0 {
			delete(rl.requests, ip)
		} else {
			rl.requests[ip] = valid
		}
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	
	// 获取该 IP 的请求记录
	times := rl.requests[ip]
	
	// 过滤掉窗口外的请求
	var validTimes []time.Time
	for _, t := range times {
		if now.Sub(t) < rl.window {
			validTimes = append(validTimes, t)
		}
	}
	
	// 检查是否超过限制
	if len(validTimes) >= rl.limit {
		return false
	}
	
	// 记录本次请求
	validTimes = append(validTimes, now)
	rl.requests[ip] = validTimes
	
	return true
}

var postLimiter = newRateLimiter(3, 1*time.Minute)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 只对 POST /api/messages 限流
		if c.Request.Method == "POST" && c.Request.URL.Path == "/api/messages" {
			ip := c.ClientIP()
			
			if !postLimiter.allow(ip) {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "rate limit exceeded, max 3 messages per minute",
				})
				c.Abort()
				return
			}
		}
		
		c.Next()
	}
}
