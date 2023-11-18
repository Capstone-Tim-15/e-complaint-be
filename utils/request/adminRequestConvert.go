package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func AdminCreateRequestToAdminDomain(request web.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request web.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Username: request.Username,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request web.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}
}

func AdminResetPasswordRequestToAdminDomain(request web.AdminResetPasswordRequest) *domain.Admin {
	return &domain.Admin{
		Password: request.NewPassword,
	}
}

func AdminDomaintoAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}
}
