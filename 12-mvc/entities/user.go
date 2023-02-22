package entities

import "time"

type UserCore struct {
	Id        uint
	Name      string
	Email     string
	Password  string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type UserResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Role    string `json:"role"`
}
