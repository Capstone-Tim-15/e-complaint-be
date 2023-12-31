package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func UserDomainToUserLoginResponse(user *domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}

func UserSchemaToUserDomain(user *schema.User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		ImageUrl: user.ImageUrl,
	}
}

func UserDomaintoUserResponse(user *domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		ImageUrl: user.ImageUrl,
	}
}

func ConvertUserResponse(users []domain.User) []web.UserResponse {
	var results []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			Id:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			ImageUrl: user.ImageUrl,
		}
		results = append(results, userResponse)
	}
	return results
}
