package web

type UserCreateRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Username string `json:"username" form:"username" validate:"required,min=1,max=16"`
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Phone    string `json:"phone" form:"phone" validate:"required,min=1,max=14"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=16"`
}

type UserLoginRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=1,max=16"`
	Password string `json:"password" form:"password" validate:"required,max=255"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Username string `json:"username" form:"username" validate:"required,min=1,max=16"`
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Phone    string `json:"phone" form:"phone" validate:"required,phone,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
}

type UserResetPasswordRequest struct {
	NewPassword        string `json:"newPassword" form:"password" validate:"min=8,max=255"`
	ConfirmNewPassword string `json:"confirmNewPassword" form:"confirmNewPassword" validate:"min=8,max=255"`
}
