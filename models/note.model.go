package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;default:uuid()" json:"id,omitempty"`
	Title     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"title,omitempty"`
	Content   string    `gorm:"type:text;not null" json:"content,omitempty"`
	Category  string    `gorm:"type:varchar(100)" json:"category,omitempty"`
	Published bool      `gorm:"default:false;not null" json:"published"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null" json:"updatedAt,omitempty"`
}

type CreateNoteSchema struct {
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Category  string `json:"category,omitempty"`
	Published bool   `json:"published,omitempty"`
}

type UpdateNoteSchema struct {
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Category  string `json:"category,omitempty"`
	Published *bool  `json:"published,omitempty"`
}
