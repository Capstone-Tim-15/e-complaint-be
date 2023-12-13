package web

type AICreateRequest struct {
	Message string `json:"message" form:"message"`
}