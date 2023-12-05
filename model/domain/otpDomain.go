package domain

type OTPUser struct {
	ID       string
	User_ID  string
	OTP      string
}

type OTPAdmin struct {
	ID       string
	Admin_ID string
	OTP      string
}
