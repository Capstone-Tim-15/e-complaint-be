package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func OTPSchemaToOTPDomain(otp *schema.OTP) *domain.OTP {
	return &domain.OTP{
		ID:      otp.ID,
		User_ID: otp.User_ID,
		OTP:     otp.OTP,
	}
}

func OTPCreateRequesttoOTPDomain(user string, OTP string) *domain.OTP {
	return &domain.OTP{
		User_ID: user,
		OTP:     OTP,
	}
}

func OTPDomaintoOTPResponse(otp *domain.OTP) web.OTPResponse {
	return web.OTPResponse{
		Id:      otp.ID,
		User_ID: otp.User_ID,
		OTP:     otp.OTP,
	}
}

func OTPDomaintoOTPCheckResponse(otp *domain.OTP) web.OTPCheckResponse {
	return web.OTPCheckResponse{
		Id:      otp.ID,
		User_ID: otp.User_ID,
		OTP:     otp.OTP,
	}
}
