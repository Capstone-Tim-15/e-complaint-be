package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
)

func OTPDomaintoOTPSchema(request *domain.OTPUser) *schema.OTPUser {
	return &schema.OTPUser{
		User_ID: request.User_ID,
		OTP:     request.OTP,
	}
}

func AdminOTPDomaintoOTPSchema(request *domain.OTPAdmin) *schema.OTPAdmin {
	return &schema.OTPAdmin{
		Admin_ID: request.Admin_ID,
		OTP:      request.OTP,
	}
}
