package web

type Message struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

type CreateRoomReq struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photoProfile"`
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
