package models

import (
	"time"
)

// Message 消息模型
type Message struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Content   string    `json:"content"`
	FromID    uint      `json:"from_id"`
	ToID      uint      `json:"to_id"`
	CreatedAt time.Time `json:"created_at"`
}
