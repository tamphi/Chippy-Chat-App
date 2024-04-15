package ws

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Controller struct {
	//todo: define struct for room controller
	board *Board
	db    *sql.DB
}

type CreateRoomRequest struct {
	//todo: define struct for create room request
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewController(b *Board, db *sql.DB) *Controller {
	return &Controller{
		board: b,
		db:    db,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//todo: specify origin to front end
		return true
	},
}

func (controllerInstance *Controller) JoinRoom(ginctx *gin.Context) {
	//todo: create new chat room
	fmt.Println("JoinRoom")
	conn, err := upgrader.Upgrade(ginctx.Writer, ginctx.Request, nil)
	if err != nil {
		ginctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatID := ginctx.Param("chatId")
	username := ginctx.Query("username")
	receiver := ginctx.Query("receiver")
	fmt.Printf("SERVER %s\n", chatID)

	//insert to db
	chechExistQuery := "SELECT content FROM messages WHERE room_name = $1"
	var existRoom int
	controllerInstance.db.QueryRowContext(ginctx, chechExistQuery, chatID).Scan(&existRoom)

	fmt.Printf("check query %d\n", existRoom)

	if _, ok := controllerInstance.board.Rooms[chatID]; !ok {
		controllerInstance.board.Rooms[chatID] = &Room{
			ID:      chatID,
			Name:    chatID,
			Clients: make(map[string]*Client),
		}
	}
	fmt.Println(username + "has joined " + chatID)
	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       username,
		Receiver: receiver,
		Username: username,
		RoomID:   chatID,
	}

	m := &Message{
		Content:   "Howdy",
		Sender:    username,
		Receiver:  receiver,
		Timestamp: time.Now(),
		Id:        uuid.NewString(),
		RoomID:    chatID,
		Username:  username,
	}

	controllerInstance.board.Register <- cl
	controllerInstance.board.Broadcast <- m
	fmt.Println(controllerInstance.board.Rooms)
	go cl.writeMessage()
	go cl.readMessage(controllerInstance)
}

func (controllerInstance *Controller) GetMessages(ginctx *gin.Context) {
	sender := ginctx.Query("username")
	receiver := ginctx.Query("receiver")
	chatId := ginctx.Param("chatId")

	db := controllerInstance.db
	query := "SELECT * FROM messages WHERE receiver=$1 OR receiver=$2 ORDER BY created_at ASC"
	rowMessages, err := db.QueryContext(context.Background(), query, sender, receiver)
	if err != nil {
		fmt.Printf("Failed at GetMessages in ws_controller: %s", err)
		return
	}
	var messages []Message
	for rowMessages.Next() {
		var id, receiver, sender, content string
		var created_at time.Time
		rowMessages.Scan(&id, &receiver, &sender, &content, &created_at)
		message := Message{
			Id:        id,
			Sender:    sender,
			Receiver:  receiver,
			Content:   content,
			Timestamp: created_at,
			RoomID:    chatId,
			Username:  sender,
		}
		messages = append(messages, message)
	}

	ginctx.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}
