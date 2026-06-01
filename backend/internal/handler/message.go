package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/service"
)

type MessageHandler struct {
	service *service.MessageService
}

func NewMessageHandler(service *service.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

type CreateMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

type MessageListResponse struct {
	Messages   interface{} `json:"messages"`
	NextCursor *int        `json:"next_cursor,omitempty"`
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var req CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ip := c.ClientIP()
	msg, err := h.service.CreateMessage(req.Content, ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, msg)
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	var cursor *int
	if cursorStr := c.Query("cursor"); cursorStr != "" {
		cursorVal, err := strconv.Atoi(cursorStr)
		if err == nil {
			cursor = &cursorVal
		}
	}

	messages, nextCursor, err := h.service.ListMessages(limit, cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, MessageListResponse{
		Messages:   messages,
		NextCursor: nextCursor,
	})
}

func (h *MessageHandler) HugMessage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message id"})
		return
	}

	msg, err := h.service.IncrementHug(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
		return
	}

	c.JSON(http.StatusOK, msg)
}

func (h *MessageHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
