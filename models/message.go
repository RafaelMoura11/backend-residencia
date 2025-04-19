package models

import (
	"time"
)

type Message struct {
	MessageID uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"agent,omitempty"`
	AgentID   uint      `gorm:"not null" json:"agent_id"`
	Role      string    `gorm:"type:enum('user','assistant');not null" json:"role"`
	Content   string    `gorm:"type:varchar(500);not null" json:"content"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`

	Agent Agent `gorm:"foreignKey:AgentID"`
}
