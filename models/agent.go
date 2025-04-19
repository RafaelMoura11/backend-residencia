package models

import "gorm.io/gorm"

type Agent struct {
	AgentID         uint   `gorm:"primaryKey" json:"agent_id"`
	Background      string `gorm:"type:varchar(500);not null" json:"background"`
	Goals           string `gorm:"type:varchar(500);not null" json:"goals"`
	Instructions    string `gorm:"type:varchar(500);not null" json:"instructions"`
	OutputIndicator string `gorm:"type:varchar(500);not null" json:"output_indicator"`
	Tasks           string `gorm:"type:varchar(500);not null" json:"tasks"`
}
