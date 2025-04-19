package models

type AgentOutputStructure struct {
	AgentOutputStructureID uint `gorm:"primaryKey;autoIncrement" json:"agent_output_structure_id"`
	AgentID                uint `gorm:"not null" json:"agent_id"`
	OutputStructureID      uint `gorm:"not null" json:"output_structure_id"`
}
