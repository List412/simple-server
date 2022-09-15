package main

import (
	"context"
	"os"
	"os/signal"
	"simple-server/internal/config"
	"simple-server/internal/server"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	err = server.Start(ctx, cfg.HttpServer)
	if err != nil {
		panic(err)
	}
}
