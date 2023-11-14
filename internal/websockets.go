package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

type Client struct {
	Conn   *websocket.Conn
	RoomId string
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
},
}

var clients = make(map[Client]bool)

func handleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	roomId := c.Param("roomId")

	if roomId == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("roomId empty")
		return
	}

	client := Client{Conn: conn, RoomId: roomId}
	clients[client] = true

	go sendPing(conn)

	for {
		messageType, _, err := conn.NextReader()
		if err != nil {
			fmt.Println("read ", err)
			delete(clients, client)
			return
		}

		if messageType == websocket.TextMessage {
			var msg Message

			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("readJson ", err)
				delete(clients, client)
				return
			}

			fmt.Println(msg)

			if msg.Type == "getFriendsOnline" {
				err := conn.WriteJSON(Message{Type: "OnlineUsers", Content: GetOnlineUsers()})
				if err != nil {
					fmt.Println(err)
					delete(clients, client)
					return
				}
			} else if msg.Type == "play" || msg.Type == "pause" {
				for client := range clients {
					err := client.Conn.WriteJSON(msg)
					if err != nil {
						fmt.Println("convert json error ", err)
						delete(clients, client)
						return
					}
				}
			}
		}
	}
}

type clientsRequest struct {
	RoomId string `json:"roomId"`
}

func GetOnlineUsers() []clientsRequest {
	var onlineClients []clientsRequest

	for client, _ := range clients {
		onlineClients = append(onlineClients, clientsRequest{RoomId: client.RoomId})
	}

	return onlineClients
}

func sendPing(conn *websocket.Conn) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
			break
		}
	}
}
