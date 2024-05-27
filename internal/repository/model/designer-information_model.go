package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DesignersInformation struct {
	ID           uuid.UUID `gorm:"primary_key;default:uuid_generate_v4();"`
	Country      string
	Company      string
	Industry     string
	DesignerType string
	DesignerName string
	ProfileUrl   string
	Status       string
	CreatedAt    time.Time `gorm:"default:now();"`
	UpdatedAt    time.Time `gorm:"default:now();"`
	DeletedAt    sql.NullTime
}
