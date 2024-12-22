package service

import (
	"context"
	"github.com/Ddarli/app/order/pkg/models"
)

type (
	Service interface {
		processOrder(ctx context.Context, request *models.OrderRequest) (*models.OrderResponse, error)
		StartConsuming(ctx context.Context)
	}
)
