package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
)

func OTPDomaintoOTPSchema(request *domain.OTP) *schema.OTPUser {
	return &schema.OTPUser{
		User_ID: request.User_ID,
		OTP:     request.OTP,
	}
}
