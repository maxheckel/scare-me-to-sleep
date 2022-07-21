package domain

import (
	"gorm.io/gorm"
	"time"
)

type Prompt struct {
	gorm.Model
	Text       string `gorm:"uniqueIndex"`
	CreatedBy  *string
	Answered   bool
	AnsweredOn *time.Time
	Priority   uint
	Responses  []Response
}

type Response struct {
	gorm.Model
	Text     string
	PromptID uint
	Votes    int32
}
