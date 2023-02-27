package delivery

import (
	"be15/clean/features/books"
	"time"
)

type BookResponse struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title" `
	UserID      uint          `json:"user_id" `
	Author      string        `json:"author" `
	Publisher   string        `json:"publisher" `
	PublishYear time.Time     `json:"publish_year" `
	User        *UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func coreToResponse(dataCore books.Core) BookResponse {
	return BookResponse{
		Id:          dataCore.Id,
		Title:       dataCore.Title,
		UserID:      dataCore.UserID,
		Author:      dataCore.Author,
		Publisher:   dataCore.Publisher,
		PublishYear: dataCore.PublishYear,
		User: &UserResponse{
			Name:  dataCore.User.Name,
			Email: dataCore.User.Email,
		},
	}
}

func listCoreToResponse(dataCore []books.Core) []BookResponse {
	var dataResponse []BookResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
