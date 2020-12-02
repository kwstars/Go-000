package dao

import (
	"github.com/kwstars/Go-000/Week02/module"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []module.Product {
	var products []module.Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByID(id uint) module.Product {
	var product module.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product module.Product) module.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product module.Product) {
	p.DB.Delete(&product)
}
