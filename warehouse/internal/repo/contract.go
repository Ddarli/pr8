package repo

import (
	"context"
	"github.com/Ddarli/utils/models"
)

type (
	Repository interface {
		GetAll(ctx context.Context) ([]*models.Product, error)
		CheckQuantity(ctx context.Context, id int, quantity int) (bool, error)
	}

	Converter interface {
		ConvertToProduct([]byte) (*models.Product, error)
	}
)
