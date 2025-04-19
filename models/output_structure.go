package models

type OutputStructure struct {
	OutputStructureID uint   `gorm:"primaryKey;autoIncrement" json:"output_structure_id"`
	Name              string `gorm:"type:varchar(255);not null" json:"name"`
	Tool              string `gorm:"type:jsonb;not null" json:"tool"`
}
