package web

import "ecomplaint/model/schema"

type LikeStatus schema.LikeStatus

type LikesCreateRequest struct {
	User_ID string     `json:"userId" form:"userId" validate:"required,min=1,max=255"`
	News_ID string     `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Status  LikeStatus `json:"status" form:"status" validate:"required"`
}

type LikesUpdateRequest struct {
	User_ID string     `json:"userId" form:"userId" validate:"required,min=1,max=255"`
	News_ID string     `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Status  LikeStatus `json:"status" form:"status" validate:"required"`
}
