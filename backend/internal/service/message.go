package service

import (
	"fmt"
	"strings"

	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/model"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/repository"
)

type MessageService struct {
	repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(content, ip string) (*model.Message, error) {
	// 验证内容长度
	content = strings.TrimSpace(content)
	if len(content) == 0 {
		return nil, fmt.Errorf("content cannot be empty")
	}
	if len(content) > 500 {
		return nil, fmt.Errorf("content exceeds 500 characters")
	}

	msg := &model.Message{
		Content:   content,
		IPAddress: ip,
	}

	if err := s.repo.Insert(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *MessageService) ListMessages(limit int, cursor *int) ([]*model.Message, *int, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	messages, err := s.repo.FindAll(limit, cursor)
	if err != nil {
		return nil, nil, err
	}

	var nextCursor *int
	if len(messages) == limit && len(messages) > 0 {
		lastID := messages[len(messages)-1].ID
		nextCursor = &lastID
	}

	return messages, nextCursor, nil
}

func (s *MessageService) IncrementHug(id int) (*model.Message, error) {
	return s.repo.UpdateHugCount(id)
}
