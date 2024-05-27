package repository

import (
	"fmt"
	"new-go-backend/internal/entity"
	"new-go-backend/internal/interface/repository"
	"new-go-backend/internal/repository/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormDIRepository struct {
	db *gorm.DB
}

func NewGormDIRepository(db *gorm.DB) repository.DIRepository {
	return &GormDIRepository{db}
}

func (repo *GormDIRepository) GetByID(id string) (*entity.DesignersInformation, error) {
	var diModel model.DesignersInformation
	if err := repo.db.First(&diModel, id).Error; err != nil {
		return nil, err
	}
	return toEntity(&diModel), nil
}

func (repo *GormDIRepository) Create(di *entity.DesignersInformation) error {
	diModel := toModel(di)
	return repo.db.Create(&diModel).Error
}

func toEntity(diModel *model.DesignersInformation) *entity.DesignersInformation {
	return &entity.DesignersInformation{
		ID:           diModel.ID.String(),
		Country:      diModel.Country,
		Company:      diModel.Company,
		Industry:     diModel.Industry,
		DesignerType: diModel.DesignerType,
		DesignerName: diModel.DesignerName,
		ProfileUrl:   diModel.ProfileUrl,
		Status:       diModel.Status,
		CreatedAt:    diModel.CreatedAt,
		UpdatedAt:    diModel.UpdatedAt,
		DeletedAt:    diModel.DeletedAt,
	}
}

func toModel(di *entity.DesignersInformation) model.DesignersInformation {
	id, err := uuid.Parse(di.ID)
	if err != nil {
		fmt.Println(err)
	}

	return model.DesignersInformation{
		ID:           id,
		Country:      di.Country,
		Company:      di.Company,
		Industry:     di.Industry,
		DesignerType: di.DesignerType,
		DesignerName: di.DesignerName,
		ProfileUrl:   di.ProfileUrl,
		Status:       di.Status,
		CreatedAt:    di.CreatedAt,
		UpdatedAt:    di.UpdatedAt,
		DeletedAt:    di.DeletedAt,
	}
}
