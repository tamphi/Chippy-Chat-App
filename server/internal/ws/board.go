package ws

import "fmt"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Board struct {
	Register   chan *Client
	Unregister chan *Client
	Rooms      map[string]*Room
	Broadcast  chan *Message
}

func NewBoard() *Board {
	return &Board{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (b *Board) Run() {
	for {
		select {
		case cl := <-b.Register:
			if _, ok := b.Rooms[cl.RoomID]; ok {
				fmt.Println("Room exists")
				r := b.Rooms[cl.RoomID]
				//check if user hasn't already joined the chat
				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-b.Unregister:
			if _, ok := b.Rooms[cl.RoomID]; ok {
				if _, ok := b.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					delete(b.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case m := <-b.Broadcast:
			if _, ok := b.Rooms[m.RoomID]; ok {

				for _, cl := range b.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}

		}

	}

}
