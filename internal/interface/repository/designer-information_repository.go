package repository

import "new-go-backend/internal/entity"

type DIRepository interface {
	GetByID(id string) (*entity.DesignersInformation, error)
	Create(di *entity.DesignersInformation) error
}
