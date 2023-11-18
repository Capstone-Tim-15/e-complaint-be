package web

type AdminReponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateAdminReponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
