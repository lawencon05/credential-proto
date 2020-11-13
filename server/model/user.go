package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Timestamp for custom type date
type Timestamp time.Time

// UsersDb model for database
type UsersDb struct {
	ID          string `gorm:"primaryKey"`
	Username    string `gorm:"unique"`
	Password    string
	Token       string
	CreatedDate Timestamp  `gorm:"type:timestamp without time zone;"`
	UpdatedDate *Timestamp `gorm:"type:timestamp without time zone;"`
}

// TableName for Alias table in database
func (UsersDb) TableName() string {
	return "tb_m_users"
}

// BeforeCreate for initiate uuid
func (base *UsersDb) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.ID = id.String()
	return nil
}
