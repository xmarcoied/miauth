package apiv1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/xmarcoied/miauth/pkg/auth"
)

// Service defines apiv1 main services
type Service struct {
	AuthService *auth.Service
}

// CreateUserCtrl creates a new user
func (s *Service) CreateUserCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
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
