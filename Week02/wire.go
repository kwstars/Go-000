//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kwstars/Go-000/Week02/dao"
	"github.com/kwstars/Go-000/Week02/service"
	"gorm.io/gorm"
)

func InitProductService(db *gorm.DB) service.ProductService {
	wire.Build(dao.ProvideProductRepostiory, service.ProvideProductService)

	return service.ProductService{}
}
