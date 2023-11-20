package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func LikeCreateRequestToLikeDomain(request web.LikesCreateRequest) *domain.Likes {
	return &domain.Likes{
		News_ID: request.News_ID,
		User_ID: request.User_ID,
		Status:  string(request.Status),
	}
}

func LikeUpdateRequestToLikeDomain(request web.LikesUpdateRequest) *domain.Likes {
	return &domain.Likes{
		News_ID: request.News_ID,
		User_ID: request.User_ID,
		Status:  string(request.Status),
	}
}

func LikeDomaintoLikeSchema(request domain.Likes) *schema.Likes {
	return &schema.Likes{
		ID:      request.ID,
		News_ID: request.News_ID,
		User_ID: request.User_ID,
		Status:  request.Status,
	}
}
