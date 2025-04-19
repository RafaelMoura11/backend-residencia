package models

type Tool struct {
	ToolID       uint   `gorm:"primaryKey;autoIncrement" json:"tool_id"`
	FunctionName string `gorm:"type:varchar(255);not null" json:"function_name"`
	Tool         string `gorm:"type:json;not null" json:"tool"` // JSON armazenado como string
}
