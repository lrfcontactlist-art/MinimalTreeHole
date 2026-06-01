package model

import "time"

type Message struct {
	ID        int       `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	HugCount  int       `json:"hug_count" db:"hug_count"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IPAddress string    `json:"-" db:"ip_address"`
}
