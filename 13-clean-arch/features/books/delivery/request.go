package delivery

import (
	"be15/clean/features/books"
	"time"
)

type BookRequest struct {
	Title       string    `json:"title" form:"title"`
	UserID      uint      `json:"user_id" form:"user_id"`
	Author      string    `json:"author" form:"author"`
	Publisher   string    `json:"publisher" form:"publisher"`
	PublishYear time.Time `json:"publish_year" form:"publish_year"`
}

func requestToCore(dataRequest BookRequest) books.Core {
	return books.Core{
		Title:       dataRequest.Title,
		UserID:      dataRequest.UserID,
		Author:      dataRequest.Author,
		Publisher:   dataRequest.Publisher,
		PublishYear: dataRequest.PublishYear,
	}
}
