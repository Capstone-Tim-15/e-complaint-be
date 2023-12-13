package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func OTPSchemaToOTPDomain(otp *schema.OTPUser) *domain.OTPUser {
	return &domain.OTPUser{
		ID:      otp.ID,
		User_ID: otp.User_ID,
		OTP:     otp.OTP,
	}
}

func AdminOTPSchemaToOTPDomain(otp *schema.OTPAdmin) *domain.OTPAdmin {
	return &domain.OTPAdmin{
		ID:      otp.ID,
		Admin_ID: otp.Admin_ID,
		OTP:     otp.OTP,
	}
}

func OTPCreateRequesttoOTPDomain(user string, OTP string) *domain.OTPUser {
	return &domain.OTPUser{
		User_ID: user,
		OTP:     OTP,
	}
}

func AdminOTPCreateRequesttoOTPDomain(admin string, OTP string) *domain.OTPAdmin {
	return &domain.OTPAdmin{
		Admin_ID: admin,
		OTP:      OTP,
	}
}

func OTPDomaintoOTPResponse(otp *domain.OTPUser) web.OTPUserResponse {
	return web.OTPUserResponse{
		Id:      otp.ID,
		User_ID: otp.User_ID,
		OTP:     otp.OTP,
	}
}

func AdminOTPDomaintoOTPResponse(otp *domain.OTPAdmin) web.OTPAdminResponse {
	return web.OTPAdminResponse{
		Id:       otp.ID,
		Admin_ID: otp.Admin_ID,
		OTP:      otp.OTP,
	}
}

func OTPDomaintoOTPCheckResponse(otp *domain.OTPUser) web.OTPUserCheckResponse {
	return web.OTPUserCheckResponse{
		Id:      otp.ID,
		User_ID: otp.User_ID,
		OTP:     otp.OTP,
	}
}

func AdminOTPDomaintoOTPCheckResponse(otp *domain.OTPAdmin) web.OTPAdminCheckResponse {
	return web.OTPAdminCheckResponse{
		Id:       otp.ID,
		Admin_ID: otp.Admin_ID,
		OTP:      otp.OTP,
	}
}
