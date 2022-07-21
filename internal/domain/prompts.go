package domain

import (
	"gorm.io/gorm"
	"time"
)

type Prompt struct {
	gorm.Model
	Text       string     `gorm:"uniqueIndex" json:"text"`
	CreatedBy  *string    `json:"created_by"`
	Answered   bool       `json:"answered"`
	AnsweredOn *time.Time `json:"answered_on"`
	Priority   uint       `json:"priority"`
	Responses  []Response `json:"responses"`
}

type Response struct {
	gorm.Model
	Text     string `json:"text"`
	PromptID uint   `json:"prompt_id"`
	Votes    int32  `json:"votes"`
}
