package web

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/xmarcoied/miauth/handlers/apiv1"
)

//Engine is the web engine wrapper
type Engine struct {
	context.Context
	httpServer *http.Server
	lock       sync.Mutex

	apiv1 *apiv1.Service
}

//New creates a new Engine
func New(apiv1 *apiv1.Service) *Engine {
	webServer := &Engine{
		Context: context.Background(),
		apiv1:   apiv1,
	}
	return webServer
}

// Run fire web Engine
func (s *Engine) Run(port int) {
	log.Info("Application is starting on port ", port)

	s.lock.Lock()
	s.httpServer = s.makeHTTPServer(port, s.routes())
	s.lock.Unlock()

	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(errors.Wrap(err, "failed to start server"))
	}
}

// Shutdown the rest server
func (s *Engine) Shutdown() error {
	s.httpServer.SetKeepAlivesEnabled(false)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s.lock.Lock()
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			log.Debugf("http server shutdown error, %s", err)
			return err
		}
	}

	log.Info("shutdown http server completed")
	s.lock.Unlock()
	return nil
}

func (s *Engine) makeHTTPServer(port int, router http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}
