package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/xmarcoied/miauth/handlers/apiv1"
	"github.com/xmarcoied/miauth/pkg/auth"
	"github.com/xmarcoied/miauth/pkg/web"
	"github.com/xmarcoied/miauth/services/storage/mongo"
)

func RunServer() int {
	svc := mongo.New(context.Background())

	api := web.New(&apiv1.Service{
		AuthService: auth.NewService(svc),
	})
	go api.Run(8080)
	return listenSignals(api, svc)
}

func listenSignals(e *web.Engine, svc *mongo.Service) int {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, append([]os.Signal{syscall.SIGTERM, syscall.SIGINT})...)
	for {
		s := <-signals
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			svc.Shutdown()
			err := e.Shutdown()
			if err != nil {
				return 1
			}
			return 0
		}
	}
}
