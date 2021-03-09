package apiv1

import (
	"net/http"

	"github.com/go-chi/render"
)

//HealthCtrl returns status true
func (s *Service) HealthCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"status": true})
	return
}
