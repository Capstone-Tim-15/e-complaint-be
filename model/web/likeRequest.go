package web

import "ecomplaint/model/schema"

type LikeStatus schema.LikeStatus

type LikesCreateRequest struct {
	User_ID string
	News_ID string     `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Status  LikeStatus `json:"status" form:"status" validate:"required"`
}

type LikesUpdateRequest struct {
	News_ID string     `json:"newsId" form:"newsId" validate:"omitempty,min=6,max=6"`
	Status  LikeStatus `json:"status" form:"status" validate:"omitempty,min=1,max=8"`
}
