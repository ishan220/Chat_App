package socket_pkg

import (
	"encoding/json"
	"fmt"

	"github.com/ishan/Chat_App/model"
)

type Pool struct {
	Clients    map[*Client]bool
	Broadcast  chan *model.Chat
	Register   chan *Client
	DeRegister chan *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *model.Chat),
		Register:   make(chan *Client),
		DeRegister: make(chan *Client),
	}
}

func (Pool *Pool) Start() {
	for {
		select {
		case chat := <-Pool.Broadcast:
			fmt.Printf("chat extracted from Broadcast %#+v \n", chat)
			//chat1 := fmt.Sprintf("{'id':%s,'from_user':%s,'to_user':%s,'message':%s,'created_at':%d,}", chat.ID, chat.FromUser, chat.ToUser, chat.Message, chat.Time)
			var response interface{}
			data := map[string]interface{}{
				"id":         chat.ID,
				"from_user":  chat.FromUser,
				"to_user":    chat.ToUser,
				"message":    chat.Message,
				"created_at": chat.Time,
			}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error from Unmarshalling")
			}
			fmt.Println("Unmarshalled Interface", response)
			for client, isValid := range Pool.Clients {
				fmt.Printf("client and is valid %+#v %b \n", client, isValid)
				if client.Username == chat.ToUser && isValid {
					fmt.Println("Writing message to client", client.Username)
					client.Conn.WriteMessage(1, jsonData)
				}
			}
			break

		case userDeregistered := <-Pool.DeRegister:
			delete(Pool.Clients, userDeregistered)
			for client, _ := range Pool.Clients {
				if client.Username != userDeregistered.Username {
					msgToClient := fmt.Sprintf("%s DeRegistered", userDeregistered.Username)
					chat := &model.Chat{FromUser: userDeregistered.Username, ToUser: client.Username, Message: msgToClient}
					client.Conn.WriteJSON(chat)

				}
			}
			break

		case userRegistered := <-Pool.Register:
			//Pool.Clients[userRegistered] = true
			for client, _ := range Pool.Clients {
				if client.Username != userRegistered.Username {
					msgToClient := fmt.Sprintf("%s Registered", userRegistered.Username)
					// data := map[string]interface{}{
					// 	"id":         strconv.FormatInt(util.RandomInt(0, 100000), 10),
					// 	"from_user":  userRegistered.Username,
					// 	"to_user":    client.Username,
					// 	"message":    msgToClient,
					// 	"created_at": time.Now().Unix(),
					// }
					// jsonData, err := json.Marshal(data)
					// if err != nil {
					// 	fmt.Println("Error from Unmarshalling")
					// }
					chat := &model.Chat{FromUser: userRegistered.Username, ToUser: client.Username, Message: msgToClient}
					//msgToClient := interface{userRegistered string}{userRegistered:userRegistered.Username}
					//fmt.Println(string(jsonData))
					//client.Conn.WriteMessage(1, []byte(jsonData))
					client.Conn.WriteJSON(chat)
				}
			}

			// 	break

		}
	}
}
