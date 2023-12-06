package websocket

type Room struct {
	ID     string              `gorm:"primaryKey" json:"id"`
	Name   string           `json:"name"`
	Client map[string]*Client `gorm:"-" json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room `gorm:"-" json:"rooms"`
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			room, ok := h.Rooms[client.RoomID]
			if !ok {
				room = &Room{
					ID:     client.ID,
					Name:   client.RoomID,
					Client: make(map[string]*Client),
				}
				h.Rooms[client.RoomID] = room
			}

			room.Client[client.ID] = client
		case client := <-h.Unregister:
			room, ok := h.Rooms[client.RoomID]
			if ok {
				if _, ok := room.Client[client.ID]; ok {
					delete(room.Client, client.ID)
					close(client.Message)
				}
			}
		case message := <-h.Broadcast:
			room, ok := h.Rooms[message.RoomID]
			if ok {
				for _, client := range room.Client {
					select {
					case client.Message <- message:
					default:
						close(client.Message)
						delete(room.Client, client.ID)
					}
				}
			}
		}
	}
}