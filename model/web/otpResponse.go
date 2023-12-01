package web

type OTPResponse struct {
	Id      string `json:"id"`
	User_ID string `json:"userId"`
	OTP     string `json:"otp"`
	Token   string `json:"token"`
}

type OTPCheckResponse struct {
	Id      string `json:"id"`
	User_ID string `json:"userId"`
	OTP     string `json:"otp"`
}
