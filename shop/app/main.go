package main

import (
	"context"
	"github.com/Ddarli/app/shop/config"
	"github.com/Ddarli/app/shop/internal/handler"
	"github.com/Ddarli/app/shop/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tokenService := service.NewTokenService(config.Key, config.TokenLifeTime)

	shopService := service.New(config.Cfg)
	go shopService.StartConsuming(ctx)

	httpHandler := handler.NewHttpHandler(shopService, tokenService)
	httpHandler.InitRouter()
}
