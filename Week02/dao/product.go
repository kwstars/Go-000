package dao

import (
	"github.com/kwstars/Go-000/Week02/module"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() ([]module.Product, error) {
	var products []module.Product
	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) FindByID(id uint) (module.Product, error) {
	var product module.Product
	if err := p.DB.First(&product, id).Error; err != nil {
		return product, errors.Wrapf(err, "Not found record, id: %d", id)
	}
	return product, nil
}
