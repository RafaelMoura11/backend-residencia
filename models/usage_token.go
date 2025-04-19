package models

type UsageToken struct {
	UsageTokenID     uint `gorm:"primaryKey;autoIncrement" json:"usage_token_id"`
	MessageID        uint `gorm:"not null" json:"message_id"`
	PromptTokens     int  `gorm:"not null" json:"prompt_tokens"`
	CompletionTokens int  `gorm:"not null" json:"completion_tokens"`
	CachedTokens     int  `gorm:"not null" json:"cached_tokens"`
	ReasoningTokens  int  `gorm:"not null" json:"reasoning_tokens"`
}
