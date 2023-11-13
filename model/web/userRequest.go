package web

type UserCreateRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,max=255"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" form:"name" validate:"min=1,max=255"`
	Email    string `json:"email" form:"email" validate:"email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"min=8,max=255"`
}

type UserResetPasswordRequest struct {
	Email              string `json:"email" form:"email" validate:"email,min=1,max=255"`
	NewPassword        string `json:"newPassword" form:"password" validate:"min=8,max=255"`
	ConfirmNewPassword string `json:"confirmNewPassword" form:"confirmNewPassword" validate:"min=8,max=255"`
}
