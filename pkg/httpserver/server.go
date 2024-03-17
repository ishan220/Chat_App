package httpserver

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/ishan/Chat_App/db/sqlc"
	"github.com/ishan/Chat_App/model"
	socket_pkg "github.com/ishan/Chat_App/pkg/websocket"
)

type HttpServer struct {
	store  *db.SQLStore
	router *gin.Engine
	Pool   *socket_pkg.Pool
}

func NewHttpServer(store *db.SQLStore, Pool *socket_pkg.Pool) *HttpServer {
	HttpServer := &HttpServer{store: store}
	HttpServer.SetUpRoutes(Pool)
	return HttpServer
}

func (server *HttpServer) SetUpRoutes(Pool *socket_pkg.Pool) {
	router := gin.Default()
	//Pool := socket_pkg.NewPool()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3000/ws"}
	router.Use(cors.New(config))
	router.POST("/create-user", server.createUser)
	router.POST("/login", server.loginUser)
	router.POST("/addContact", server.AddContact)
	router.GET("/getContacts", server.GetContacts)
	router.GET("/ChatBtwnUsers", server.ChatBtwnUsers)
	//wsRoute := router.Group("/").Use(WebSocketMiddleWare(Pool, server))
	//wsRoute.GET("/ws", WebSocketHandler)
	router.GET("/ws", server.WebSocketHandler)
	server.Pool = Pool
	server.router = router
}

func (server *HttpServer) Run() error {
	return server.router.Run("localhost:8080")
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func WebSocketMiddleWare(Pool *socket_pkg.Pool, Server *HttpServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Pool", Pool)
		ctx.Set("Server", Server)
		ctx.Next()
	}
}

func (Server *HttpServer) WebSocketHandler(ctx *gin.Context) {

	conn, err := socket_pkg.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	fmt.Println("Wbsocket Connection sucessful")

	client := &socket_pkg.Client{
		Conn: conn,
		Pool: Server.Pool, //this key is set in middleware
	}

	defer func() {
		conn.Close()
		client.Pool.DeRegister <- client
		delete(client.Pool.Clients, client)
	}()

	for {
		fmt.Println("start reading the message:")
		_, p, err := client.Conn.ReadMessage()
		fmt.Println("Read the message:", string(p))
		if err != nil {
			return
		}

		fmt.Println("Message received at Golang server", string(p))
		m := &socket_pkg.Message{}
		err1 := json.Unmarshal(p, m)

		if err1 != nil {
			fmt.Println("Error", err1)
			return
		}

		if m.Type == "Bootup" {
			fmt.Println("Got Bootup Message")
			fmt.Println("Client Registered", client)
			client.Username = m.User
			client.Pool.Register <- client
			client.Pool.Clients[client] = true
		} else {
			fmt.Println("Non Bootup Message")
			chat := &model.Chat{}
			chat.ID = m.Chat.ID
			chat.FromUser = m.Chat.FromUser
			chat.ToUser = m.Chat.ToUser
			chat.Message = m.Chat.Message
			chat.Time = time.Now().Unix()

			ctx.Set("chat", chat)
			Server.CreateChat(ctx)
			client.Pool.Broadcast <- chat

		}
	}

}
