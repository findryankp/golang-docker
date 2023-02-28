package delivery

import "be15/clean/features/users"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func requestToCore(dataRequest UserRequest) users.Core {
	return users.Core{
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Address:  dataRequest.Address,
		Role:     dataRequest.Role,
	}
}
