package repository

import (
	"gorm.io/gorm"

	"github.com/d60-Lab/gin-template/internal/model"
)

type PracticeRouteRepository interface {
	List(category string) ([]model.PracticeRoute, error)
	GetByID(id int64) (*model.PracticeRoute, error)
	Create(route *model.PracticeRoute) error
	Update(route *model.PracticeRoute) error
	Delete(id int64) error
}

type practiceRouteRepository struct {
	db *gorm.DB
}

func NewPracticeRouteRepository(db *gorm.DB) PracticeRouteRepository {
	return &practiceRouteRepository{db: db}
}

func (r *practiceRouteRepository) List(category string) ([]model.PracticeRoute, error) {
	var routes []model.PracticeRoute
	query := r.db.Where("is_deleted = ?", false)

	if category != "" {
		query = query.Where("name = ?", category)
	}

	err := query.Order("sort ASC").Preload("Phases", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort ASC")
	}).Find(&routes).Error

	return routes, err
}

func (r *practiceRouteRepository) GetByID(id int64) (*model.PracticeRoute, error) {
	var route model.PracticeRoute
	err := r.db.Where("id = ? AND is_deleted = ?", id, false).
		Preload("Phases", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort ASC")
		}).First(&route).Error

	if err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *practiceRouteRepository) Create(route *model.PracticeRoute) error {
	return r.db.Create(route).Error
}

func (r *practiceRouteRepository) Update(route *model.PracticeRoute) error {
	return r.db.Save(route).Error
}

func (r *practiceRouteRepository) Delete(id int64) error {
	return r.db.Model(&model.PracticeRoute{}).Where("id = ?", id).Update("is_deleted", true).Error
}
