package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func LikesSchemaToLikesDomain(likes *schema.Likes) *domain.Likes {
	return &domain.Likes{
		ID:      likes.ID,
		News_ID: likes.News_ID,
		User_ID: likes.User_ID,
		Status:  string(likes.Status),
	}
}

func LikesDomainToLikesResponse(likes *domain.Likes) web.LikesResponse {
	return web.LikesResponse{
		ID:      likes.ID,
		News_ID: likes.News_ID,
		User_ID: likes.User_ID,
		Status:  likes.Status,
	}
}

func ConvertLikesResponse(likes []domain.Likes) []web.LikesResponse {
	var results []web.LikesResponse
	for _, l := range likes {
		likesResponse := web.LikesResponse{
			ID:      l.ID,
			News_ID: l.News_ID,
			User_ID: l.User_ID,
			Status:  l.Status,
		}
		results = append(results, likesResponse)
	}
	return results
}
