package apiv1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/xmarcoied/miauth/handlers"
	"github.com/xmarcoied/miauth/pkg"
	"github.com/xmarcoied/miauth/pkg/auth"
)

// Service defines apiv1 main services
type Service struct {
	AuthService *auth.Service
}

// CreateUserCtrl creates a new user
func (s *Service) CreateUserCtrl(w http.ResponseWriter, r *http.Request) {
	var entity auth.CreateUserRequest
	if err := handlers.BindTo(w, r, &entity); err != nil {
		return
	}
	if err := handlers.Validation(w, r, entity); err != nil {
		return
	}

	_, err := s.AuthService.CreateUser(r.Context(), entity)
	if err != nil {
		handlers.RenderJSONError(w, r, http.StatusBadRequest, &pkg.Error{
			Code: pkg.ErrInternal,
			Msg:  "",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	return
}

// LoginCtrl tries to login with username and password
func (s *Service) LoginCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	return
}

// UpdateUserCtrl updates user's info
func (s *Service) UpdateUserCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	return
}

// ResetPasswordCtrl resets user's password
func (s *Service) ResetPasswordCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	return
}

// ChangePasswordCtrl tries to change user's password
func (s *Service) ChangePasswordCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	return
}
