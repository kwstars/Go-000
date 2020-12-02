package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kwstars/Go-000/Week02/dao"
	"github.com/kwstars/Go-000/Week02/module"
	"log"
	"net/http"
	"strconv"
)

type ProductService struct {
	ProductRepository dao.ProductRepository
}

func ProvideProductService(p dao.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll(c *gin.Context) {
	products := p.ProductRepository.FindAll()
	c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

func (p *ProductService) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductRepository.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

func (p *ProductService) Create(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.ProductRepository.Save(ToProduct(productDTO))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createdProduct)})
}

func (p *ProductService) Update(c *gin.Context) {
	var productDTO ProductDTO
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductRepository.FindByID(uint(id))
	if product == (module.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductRepository.Save(product)

	c.Status(http.StatusOK)
}

func (p *ProductService) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductRepository.FindByID(uint(id))
	if product == (module.Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductRepository.Delete(product)

	c.Status(http.StatusOK)
}

type ProductDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price,string"`
}

func ToProduct(productDTO ProductDTO) module.Product {
	return module.Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product module.Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []module.Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
