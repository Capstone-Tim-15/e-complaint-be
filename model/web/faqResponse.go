package web

type FaqResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category_ID string `json:"categoryId"`
}
