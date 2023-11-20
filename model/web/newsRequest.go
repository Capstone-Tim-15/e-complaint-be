package web

import (
	"ecomplaint/model/domain"
	"time"
)

type Feedback domain.Feedback
type Likes domain.Likes

type NewsCreateRequest struct {
	Admin_ID string     `json:"adminId" form:"adminId" validate:"required,min=1,max=255"`
	Title    string     `json:"title" form:"title" validate:"required,min=1,max=255"`
	Content  string     `json:"content" form:"content" validate:"required,min=1,max=255"`
	Date     time.Time  `json:"date" form:"date" `
	Feedback []Feedback `json:"feedback" form:"feedback" validate:"required"`
	Likes    []Likes    `json:"likes" form:"likes" validate:"required"`
}

type NewsUpdateRequest struct {
	Admin_ID string     `json:"adminId" form:"adminId" validate:"required,min=1,max=255"`
	Title    string     `json:"title" form:"title" validate:"required,min=1,max=255"`
	Content  string     `json:"content" form:"content" validate:"required,min=1,max=255"`
	Date     time.Time  `json:"date" form:"date" `
	Feedback []Feedback `json:"feedback" form:"feedback" validate:"required"`
	Likes    []Likes    `json:"likes" form:"likes" validate:"required"`
}
