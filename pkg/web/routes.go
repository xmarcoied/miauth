package web

import "github.com/go-chi/chi"

func (s *Engine) routes() chi.Router {
	router := chi.NewRouter()

	//routes
	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Get("/_health", s.apiv1.HealthCtrl)
		rapi.Route("/user", func(ruser chi.Router) {
			ruser.Post("/", s.apiv1.CreateUserCtrl)
			ruser.Post("/login", s.apiv1.LoginCtrl)

			ruser.Route("/{id}", func(ri chi.Router) {
				ri.Put("/", s.apiv1.UpdateUserCtrl)
				ri.Post("/change_password", s.apiv1.ChangePasswordCtrl)
				ri.Post("/reset_password", s.apiv1.ResetPasswordCtrl)
			})
		})
	})
	return router
}
