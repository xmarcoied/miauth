package web

import "github.com/go-chi/chi"

func (s *Engine) routes() chi.Router {
	router := chi.NewRouter()

	//routes
	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Get("/_health", s.apiv1.HealthCtrl)
		rapi.Route("/user", func(ruser chi.Router) {
			ruser.Post("/", nil)
			ruser.Post("/login", nil)

			ruser.Route("/{id}", func(ri chi.Router) {
				ri.Put("/", nil)
				ri.Post("/change_password", nil)
				ri.Post("/reset_password", nil)
			})
		})
	})
	return router
}
