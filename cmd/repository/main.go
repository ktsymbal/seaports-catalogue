package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/internal/repository"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"
	"github.com/kateGlebova/seaports-catalogue/pkg/storage/inmem"

	"github.com/kateGlebova/seaports-catalogue/pkg/shutdown"
)

func main() {
	storage := inmem.NewRepository()
	grpcSvc := proto.NewRepositoryService(storage)
	portDomainSvc := repository.NewPortDomainService(grpcSvc, "9090")

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, shutdown.GracefulShutdownSignals...)
	go shutdown.SignalHandle(signalChan, exitChan, portDomainSvc.Stop)
	go portDomainSvc.Run()
	code := <-exitChan
	os.Exit(code)
}
