package auth

// CreateUserRequest
//
// swagger:model
type CreateUserRequest struct {
	// required: true
	Password string `json:"password" validate:"required"`
	// required: true
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	// required: true
	Username  string `json:"username" validate:"required"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// LoginRequest
//
// swagger:model
type LoginRequest struct {
	// required: true
	Username string `json:"username" validate:"required"`
	// required: true
	Password string `json:"password" validate:"required"`
}
