package res

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func UserDomainToUserLoginResponse(user *domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func UserSchemaToUserDomain(user *schema.User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UserDomaintoUserResponse(user *domain.User) web.UserResponse {
	return web.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ConvertUserResponse(users []domain.User) []web.UserResponse {
	var results []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		results = append(results, userResponse)
	}
	return results
}
