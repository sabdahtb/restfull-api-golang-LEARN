package dto

// LoginDTO is a model when user is login
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}
