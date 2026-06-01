package repository

import (
	"database/sql"
	"fmt"

	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/model"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Insert(msg *model.Message) error {
	query := `
		INSERT INTO messages (content, ip_address)
		VALUES ($1, $2)
		RETURNING id, hug_count, created_at
	`
	err := r.db.QueryRow(query, msg.Content, msg.IPAddress).Scan(&msg.ID, &msg.HugCount, &msg.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}
	return nil
}

func (r *MessageRepository) FindAll(limit int, cursor *int) ([]*model.Message, error) {
	var query string
	var args []interface{}

	if cursor != nil {
		query = `
			SELECT id, content, hug_count, created_at
			FROM messages
			WHERE id < $1
			ORDER BY created_at DESC
			LIMIT $2
		`
		args = []interface{}{*cursor, limit}
	} else {
		query = `
			SELECT id, content, hug_count, created_at
			FROM messages
			ORDER BY created_at DESC
			LIMIT $1
		`
		args = []interface{}{limit}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query messages: %w", err)
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		msg := &model.Message{}
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.HugCount, &msg.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (r *MessageRepository) UpdateHugCount(id int) (*model.Message, error) {
	query := `
		UPDATE messages
		SET hug_count = hug_count + 1
		WHERE id = $1
		RETURNING id, content, hug_count, created_at
	`
	msg := &model.Message{}
	err := r.db.QueryRow(query, id).Scan(&msg.ID, &msg.Content, &msg.HugCount, &msg.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update hug count: %w", err)
	}
	return msg, nil
}
