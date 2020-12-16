package dao

import (
	"github.com/google/wire"
	"github.com/kwstars/Go-000/Week04/internal/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var ProvierSets = wire.NewSet(New)

type ProductRepository struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	if err := p.DB.First(&product, id).Error; err != nil {
		return product, errors.Wrapf(err, "Not found record, id: %d", id)
	}
	return product, nil
}
