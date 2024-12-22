package main

import (
	"context"
	"github.com/Ddarli/app/warehouse/config"
	"github.com/Ddarli/app/warehouse/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	warehouseService := service.New(config.Cfg)
	warehouseService.StartConsuming(ctx)
}
