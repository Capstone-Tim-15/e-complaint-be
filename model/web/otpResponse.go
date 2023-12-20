package web

type OTPUserResponse struct {
	Id      string `json:"id"`
	User_ID string `json:"userId"`
	Token   string `json:"token"`
}

type OTPAdminResponse struct {
	Id       string `json:"id"`
	Admin_ID string `json:"adminId"`
	Token    string `json:"token"`
}

type OTPUserCheckResponse struct {
	Id       string `json:"id"`
	User_ID  string `json:"userId"`
	OTP      string `json:"otp"`
}

type OTPAdminCheckResponse struct {
	Id       string `json:"id"`
	Admin_ID string `json:"adminId"`
	OTP      string `json:"otp"`
}
