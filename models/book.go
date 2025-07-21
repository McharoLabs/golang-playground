package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string         `json:"title" gorm:"not null"`
	Author    string         `json:"author" gorm:"not null"`
	Year      int            `json:"year"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

func (Book) TableName() string {
	return "books"
}
