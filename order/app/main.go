package main

import (
	"context"
	"github.com/Ddarli/app/order/config"
	"github.com/Ddarli/app/order/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	orderService := service.New(config.Cfg)
	orderService.StartConsuming(ctx)
}
