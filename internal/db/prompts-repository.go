package db

import (
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"gorm.io/gorm"
)

type PromptsRepository interface {
	GetToday() (*domain.Prompt, error)
	GetDay(daysBack int) (*domain.Prompt, error)
}

type Prompts struct {
	DB *gorm.DB
}

func (p Prompts) GetDay(daysBack int) (*domain.Prompt, error) {

	panic("implement me")
}

func NewPromptsRepository(db *gorm.DB) PromptsRepository {
	return Prompts{DB: db}
}

func (p Prompts) GetToday() (*domain.Prompt, error) {
	nextPrompt := &domain.Prompt{}
	err := p.DB.Preload("Responses").Last(&nextPrompt, "answered = true").Order("answered_on ASC").Error
	return nextPrompt, err
}
