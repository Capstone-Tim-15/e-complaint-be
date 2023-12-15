package web

import (
	"time"
)

type NewsCreateRequest struct {
	Admin_ID    string
	Category_ID string    `json:"categoryId" form:"categoryId" validate:"min=6,max=6"`
	Title       string    `json:"title" form:"title" validate:"required,min=1,max=255"`
	Content     string    `json:"content" form:"content" validate:"required,min=1,max=255"`
	Date        time.Time `json:"date" form:"date"`
	ImageUrl    string    `json:"imageUrl" form:"imageUrl" validate:"omitempty,min=1,max=255"`
}

type NewsUpdateRequest struct {
	Category_ID string    `json:"categoryId" form:"categoryId" validate:"omitempty,min=6,max=6"`
	Title       string    `json:"title" form:"title" validate:"omitempty,min=1,max=255"`
	Content     string    `json:"content" form:"content" validate:"omitempty,min=1,max=255"`
	Date        time.Time `json:"date" form:"date" `
	ImageUrl    string    `json:"imageUrl" form:"imageUrl" validate:"omitempty,min=1,max=255"`
}
