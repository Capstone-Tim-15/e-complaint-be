package web

type LikesCreateRequest struct {
	User_ID string     `json:"userId" form:"userId" validate:"required,min=1,max=255"`
	News_ID string     `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Status  likeStatus `json:"status" form:"status" validate:"required,min=1,max=255"`
}

type LikesUpdateRequest struct {
	User_ID string     `json:"userId" form:"userId" validate:"required,min=1,max=255"`
	News_ID string     `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Status  likeStatus `json:"status" form:"status" validate:"required,min=1,max=255"`
}

type likeStatus string

const (
	LIKE   likeStatus = "LIKE"
	UNLIKE likeStatus = "DISLIKE"
)
