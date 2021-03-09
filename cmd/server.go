package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/xmarcoied/miauth/handlers/apiv1"
	"github.com/xmarcoied/miauth/pkg/auth"
	"github.com/xmarcoied/miauth/pkg/web"
	"github.com/xmarcoied/miauth/services/storage"
)

func RunServer() int {

	api := web.New(&apiv1.Service{
		AuthService: auth.NewService(&storage.MockedUser{}),
	})
	go api.Run(8080)
	return listenSignals(api)
}

func listenSignals(e *web.Engine) int {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, append([]os.Signal{syscall.SIGTERM, syscall.SIGINT})...)
	for {
		s := <-signals
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			err := e.Shutdown()
			if err != nil {
				return 1
			}
			return 0
		}
	}
}
