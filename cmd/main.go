package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/Amartha/go-dlq-retrier/internal/handler"
	"bitbucket.org/Amartha/go-dlq-retrier/internal/handler/health"
	server "bitbucket.org/Amartha/go-dlq-retrier/internal/http"
)

func main() {

	appContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	srv := server.New(handler.New(health.NewGet()))
	srv.Start(appContext)

	<-appContext.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	srv.Close(ctx)

}
