package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id" validate:"omitempty,uuid"`
	FirstName    string         `json:"first_name" gorm:"not null" validate:"required,min=2,max=50,alpha"`
	LastName     string         `json:"last_name" gorm:"not null" validate:"required,min=2,max=50,alpha"`
	Email        string         `json:"email" gorm:"not null;unique" validate:"required,email,max=100"`
	PhoneNumber  string         `json:"phone_number" gorm:"not null;unique" validate:"required,e164"`
	Token        string         `json:"token"`
	UserType     string         `json:"user_type" gorm:"not null" validate:"required,oneof=ADMIN USER"`
	RefreshToken string         `json:"refresh_token"`
	Password     string         `json:"password" gorm:"not null" validate:"required,min=6,max=100"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName sets the table name explicitly
func (User) TableName() string {
	return "users"
}

// BeforeCreate sets a UUID before creating a record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

// Validate runs validator on User struct
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
