package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
)

func OTPDomaintoOTPSchema(request *domain.OTP) *schema.OTP {
	return &schema.OTP{
		User_ID: request.User_ID,
		OTP:     request.OTP,
	}
}
