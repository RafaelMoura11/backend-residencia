package models

import (
	"time"
)

type Message struct {
	MessageID uint      `gorm:"primaryKey;autoIncrement" json:"message_id"`
	AgentID   uint      `gorm:"not null" json:"agent_id"`
	Role      string    `gorm:"type:varchar(20);not null" json:"role"`
	Content   string    `gorm:"type:varchar(500);not null" json:"content"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
}
