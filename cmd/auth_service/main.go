package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"snakealive/m/internal/services/auth/config"
	"snakealive/m/internal/services/auth/setup"
)

func main() {
	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		return
	}

	logger := cfg.Logger.Sugar()
	server, cancel, err := setup.SetupServer(cfg)
	if err != nil {
		logger.Fatal("msg", "failed to setup server", "error", err)
		return
	}

	go func() {
		logger.Info("msg", "starting grpc server", "port", cfg.GRPCPort)
		lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
		if err != nil {
			logger.Error("msg", "grpc server listen", "err", err)
			os.Exit(1)
		}
		logger.Info("msg", "grpc server listener started")

		if err := server.Serve(lis); err != nil {
			logger.Error("msg", "grpc server run failuer", "err", err)
			os.Exit(1)
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	logger.Info("auth service started ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		logger.Info("msg", "received signal, exiting", "signal", sig)
		server.GracefulStop()
		cfg.Cancel()
		cancel()

		logger.Info("msg", " goodbye")
	}(<-c)
}
