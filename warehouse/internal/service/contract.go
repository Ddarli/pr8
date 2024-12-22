package service

import (
	"context"
)

type (
	Service interface {
		StartConsuming(ctx context.Context)
		GetProducts(ctx context.Context)
	}
)
