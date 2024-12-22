package service

import (
	"context"
	pkg "github.com/Ddarli/app/shop/pkg/models"
	"github.com/Ddarli/utils/models"
	"github.com/golang-jwt/jwt/v5"
)

type (
	Service interface {
		GetAll(ctx context.Context) ([]*models.Product, error)
		ProcessOrder(ctx context.Context, request pkg.OrderRequest) (*pkg.OrderResponse, error)
		StartConsuming(ctx context.Context)
	}

	TokenService interface {
		GenerateAccessToken(userID string) (string, error)
		ValidateToken(tokenString string) (*jwt.Token, error)
	}
)
