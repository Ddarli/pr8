package pkg

import (
	"github.com/Ddarli/utils/models"
	"strconv"
)

type (
	Product struct {
		ID       int
		Name     string
		Price    float64
		Quantity int
	}

	OrderRequest struct {
		ID       int
		Name     string
		Quantity int
	}
)

type (
	Converter interface {
		ToProto(products []Product) []*models.Product
	}
	ProductConverter struct{}
)

func (p *ProductConverter) ToProto(products []Product) []*models.Product {
	var protoProducts []*models.Product

	for _, product := range products {
		protoProducts = append(protoProducts, &models.Product{
			Id:    strconv.Itoa(product.ID),
			Name:  product.Name,
			Price: float32(product.Price),
		})
	}

	return protoProducts
}
