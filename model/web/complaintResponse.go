package web

type ComplaintResponse struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	URL        string `json:"url"`
}
