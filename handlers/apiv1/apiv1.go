package apiv1

import "github.com/xmarcoied/miauth/pkg/auth"

// Service defines apiv1 main services
type Service struct {
	AuthService *auth.Service
}
