package apiv1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/xmarcoied/miauth/handlers"
	"github.com/xmarcoied/miauth/pkg"
	"github.com/xmarcoied/miauth/pkg/auth"
	"github.com/xmarcoied/miauth/services/storage"
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
		if errors.Is(err, storage.ErrAlreadyExist) {
			handlers.RenderJSONError(w, r, http.StatusConflict, &pkg.Error{
				Code: pkg.ErrInternal,
				Msg:  "username already exist",
			})
			return
		}
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
	var entity auth.LoginRequest
	if err := handlers.BindTo(w, r, &entity); err != nil {
		return
	}
	if err := handlers.Validation(w, r, entity); err != nil {
		return
	}

	err := s.AuthService.Login(r.Context(), entity)
	if err != nil {
		handlers.RenderJSONError(w, r, http.StatusBadRequest, &pkg.Error{
			Code: pkg.ErrInternal,
			Msg:  err.Error(),
		})
		return
	}

	render.Status(r, http.StatusOK)
	return
}

// UpdateUserCtrl updates user's info
func (s *Service) UpdateUserCtrl(w http.ResponseWriter, r *http.Request) {
	var entity auth.UpdateUserRequest
	if err := handlers.BindTo(w, r, &entity); err != nil {
		return
	}
	if err := handlers.Validation(w, r, entity); err != nil {
		return
	}

	username := chi.URLParam(r, "username")

	err := s.AuthService.UpdateUser(r.Context(), username, entity)
	if err != nil {
		handlers.RenderJSONError(w, r, http.StatusBadRequest, &pkg.Error{
			Code: pkg.ErrInternal,
			Msg:  err.Error(),
		})
		return
	}

	render.Status(r, http.StatusOK)
	return
}

// ResetPasswordCtrl resets user's password
func (s *Service) ResetPasswordCtrl(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	password, err := s.AuthService.ResetPassword(r.Context(), username)
	if err != nil {
		handlers.RenderJSONError(w, r, http.StatusBadRequest, &pkg.Error{
			Code: pkg.ErrInternal,
			Msg:  err.Error(),
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{"new_password": password})
	return
}

// ChangePasswordCtrl tries to change user's password
func (s *Service) ChangePasswordCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	return
}
