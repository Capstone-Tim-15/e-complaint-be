package req

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.User {
	return &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserLoginRequestToUserDomain(request web.UserLoginRequest) *domain.User {
	return &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserUpdateRequestToUserDomain(request web.UserUpdateRequest) *domain.User {
	return &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserResetPasswordRequestToUserDomain(request web.UserResetPasswordRequest) *domain.User {
	return &domain.User{
		Password: request.NewPassword,
	}
}

func UserDomaintoUserSchema(request domain.User) *schema.User {
	return &schema.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
