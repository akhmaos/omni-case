package main

import (
	"fmt"

	"github.com/akhmaos/omni-case/internal/api"
	"github.com/akhmaos/omni-case/internal/service"
	"github.com/akhmaos/omni-case/internal/usecase"
)

func main() {
	// в аргументах передаем максимальное количество итемов которое может обрабатывать внешний сервис
	srv := service.NewService(5)

	client := usecase.NewClient(srv)

	authUIServer, err := api.NewServer(
		client,
		"0.0.0.0",
		8080,
	)
	if err != nil {
		panic(fmt.Sprintf("error with init server %s", err))
	}

	if err = authUIServer.Serve(); err != nil {
		panic(fmt.Sprintf("server stopped %s", err))
	}
}
