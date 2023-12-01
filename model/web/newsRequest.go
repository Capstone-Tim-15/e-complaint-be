package web

import (
	"ecomplaint/model/domain"
	"time"
)

type NewsCreateRequest struct {
	Admin_ID string            `json:"adminId" form:"adminId" validate:"required,min=1,max=255"`
	Title    string            `json:"title" form:"title" validate:"required,min=1,max=255"`
	Content  string            `json:"content" form:"content" validate:"required,min=1,max=255"`
	Date     time.Time         `json:"date" form:"date" `
	Feedback []domain.Feedback `json:"feedback"`
	Likes    []domain.Likes    `json:"likes"`
}

type NewsUpdateRequest struct {
	Admin_ID string            `json:"adminId" form:"adminId" validate:"required,min=1,max=255"`
	Title    string            `json:"title" form:"title" validate:"required,min=1,max=255"`
	Content  string            `json:"content" form:"content" validate:"required,min=1,max=255"`
	Date     time.Time         `json:"date" form:"date" `
	Feedback []domain.Feedback `json:"feedback"`
	Likes    []domain.Likes    `json:"likes"`
}
