package web

type OTPCreateRequest struct {
	Email string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
}

type OTPCheckRequest struct {
	OTP string `json:"otp" form:"otp" validate:"required,min=6,max=6"`
}
