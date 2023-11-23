package web

type CategoryResponse struct {
	Id           string                `json:"id"`
	CategoryName string                `json:"CategoryName"`
	FAQ          []FaqCategoryResponse `json:"faq"`
}

type CategoryCreateResponse struct {
	Id           string `json:"id"`
	CategoryName string `json:"CategoryName"`
}
