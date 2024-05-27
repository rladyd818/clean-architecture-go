package entity

import (
	"database/sql"
	"time"
)

type DesignersInformation struct {
	ID           string
	Country      string
	Company      string
	Industry     string
	DesignerType string
	DesignerName string
	ProfileUrl   string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}
