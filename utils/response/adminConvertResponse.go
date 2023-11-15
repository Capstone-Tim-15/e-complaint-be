package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func AdminDomainToAdminLoginResponse(admin *domain.Admin) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Name:  admin.Name,
		Email: admin.Email,
	}
}

func AdminSchemaToAdminDomain(admin *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:       admin.ID,
		Name:     admin.Name,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Password: admin.Password,
	}
}

func AdminDomaintoAdminResponse(admin *domain.Admin) web.AdminReponse {
	return web.AdminReponse{
		Id:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Phone: admin.Phone,
	}
}

func ConvertAdminResponse(admins []domain.Admin) []web.AdminReponse {
	var results []web.AdminReponse
	for _, admin := range admins {
		adminResponse := web.AdminReponse{
			Id:    admin.ID,
			Name:  admin.Name,
			Email: admin.Email,
			Phone: admin.Phone,
		}
		results = append(results, adminResponse)
	}
	return results
}
