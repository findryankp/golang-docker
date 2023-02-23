package entities

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
