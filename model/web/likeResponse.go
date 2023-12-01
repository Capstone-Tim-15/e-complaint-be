package web

type LikesResponse struct {
	ID      string `json:"id"`
	User_ID string `json:"userId"`
	News_ID string `json:"newsId"`
	Status  string `json:"status"`
}
