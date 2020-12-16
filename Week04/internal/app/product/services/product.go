package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kwstars/Go-000/Week04/internal/app/product/dao"
	"github.com/kwstars/Go-000/Week04/internal/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

var ProviderSet = wire.NewSet(New)

type ProductService struct {
	ProductRepository *dao.ProductRepository
}

func New(p *dao.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll(c *gin.Context) {
	products, err := p.ProductRepository.FindAll()
	switch {
	//不会返回ErrRecordNotFound错误
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusOK, gin.H{"code": "1", "content": "没有找到记录"})
	case err != nil:
		log.Printf("FindALl failed, %+v", err)
		c.JSON(http.StatusOK, gin.H{"code": "1", "content": "未知错误"})
	default:
		c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
	}
}

func (p *ProductService) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.ProductRepository.FindByID(uint(id))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusOK, gin.H{"code": "1", "content": "没有找到记录"})
	case err != nil:
		log.Printf("FindALl failed, %+v", err)
		c.JSON(http.StatusOK, gin.H{"code": "1", "content": "未知错误"})
	default:
		c.JSON(http.StatusOK, gin.H{"products": ToProductDTO(product)})
	}
}

type ProductDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price,string"`
}

func ToProductDTO(product models.Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []models.Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
