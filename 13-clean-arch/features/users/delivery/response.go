package delivery

import "be15/clean/features/users"

type UserResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Role    string `json:"role"`
}

func coreToResponse(dataCore users.Core) UserResponse {
	return UserResponse{
		Id:      dataCore.Id,
		Name:    dataCore.Name,
		Email:   dataCore.Email,
		Address: dataCore.Address,
		Role:    dataCore.Role,
	}
}

func listCoreToResponse(dataCore []users.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
