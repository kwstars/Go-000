package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kwstars/Go-000/Week04/internal/app/product/services"
	"github.com/kwstars/Go-000/Week04/internal/pkg/transports/http"
)

var ProviderSet = wire.NewSet(CreateInitControllersFn, New)

func CreateInitControllersFn(pc *ProductsController) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/products", pc.service.FindAll)
		r.GET("/products/:id", pc.service.FindByID)
	}
}

func New(s *services.ProductService) *ProductsController {
	return &ProductsController{
		service: s,
	}
}

type ProductsController struct {
	service *services.ProductService
}
