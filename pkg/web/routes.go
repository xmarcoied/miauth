// Auth microservice
//
// This documentation describes miauth APIs
//
//     Schemes: https
//	   Host: https://NONE
//     BasePath: /api/v1
//     Version: 1.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/xmarcoied/miauth/pkg"
)

func (s *Engine) routes() chi.Router {
	router := chi.NewRouter()

	router.Use(UseRequestID("Request_ID"))

	//routes
	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Get("/_health", s.apiv1.HealthCtrl)
		rapi.Route("/user", func(ruser chi.Router) {
			ruser.Post("/", s.apiv1.CreateUserCtrl)
			ruser.Post("/login", s.apiv1.LoginCtrl)

			ruser.Route("/{username}", func(ri chi.Router) {
				ri.Put("/", s.apiv1.UpdateUserCtrl)
				ri.Post("/change_password", s.apiv1.ChangePasswordCtrl)
				ri.Post("/reset_password", s.apiv1.ResetPasswordCtrl)
			})
		})
	})
	return router
}

// UseRequestID middleware adds request-id to the context
func UseRequestID(headerName string) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var reqID string
			if reqID = r.Header.Get(headerName); reqID == "" {
				// first occurrence
				reqID = pkg.GenerateNewStringUUID()
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, pkg.RequestID, reqID)

			// setup logger
			requestLogger := log.WithField(pkg.RequestID, reqID)
			ctx = context.WithValue(ctx, pkg.RequestID, requestLogger)

			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
	return f
}
