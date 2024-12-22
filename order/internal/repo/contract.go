package repo

import (
	"context"
	"github.com/Ddarli/app/order/pkg/models"
	utils "github.com/Ddarli/utils/models"
)

type (
	Repository interface {
		SaveOrder(ctx context.Context, order *models.Order) (bool, error)
	}

	Converter interface {
		ConvertToProduct([]byte) (*utils.Product, error)
	}
)
