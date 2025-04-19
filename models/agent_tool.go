package models

type AgentTool struct {
	AgentToolID uint `gorm:"primaryKey;autoIncrement" json:"agent_tool_id"`
	AgentID     uint `gorm:"not null" json:"agent_id"`
	ToolID      uint `gorm:"not null" json:"tool_id"`
}
