package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/ishan/Chat_App/db/sqlc"
	"github.com/ishan/Chat_App/model"
	socket_pkg "github.com/ishan/Chat_App/pkg/websocket"
)

// var ginLambda *ginadapter.GinLambda

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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000",
		"http://localhost:3001", "http://localhost:3000/ws",
		"https://chat-on-go.netlify.app",
		"https://chat-app-delta-five.vercel.app/",
		"https://go-chat-app.surge.sh",
		"https://chatapp-production-0da2.up.railway.app/"}
	router.Use(cors.New(config))
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Status": "Success"})
		return
	})

	router.POST("/create-user", server.createUser)
	router.POST("/login", server.loginUser)
	router.POST("/addContact", server.AddContact)
	router.GET("/getContacts", server.GetContacts)
	router.GET("/ChatBtwnUsers", server.ChatBtwnUsers)
	router.GET("/ws", server.WebSocketHandler)
	server.Pool = Pool
	server.router = router
	//ginLambda = ginadapter.New(router)
}

func (server *HttpServer) Run() error {
	//return server.router.Run("chat-app-delta-five.vercel.app:8080")
	return server.router.Run("0.0.0.0:8080")
	//return server.router.Run("https://chat-on-the-go.up.railway.app/")
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

// func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	// If no name is provided in the HTTP request body, throw an error
// 	return ginLambda.ProxyWithContext(ctx, req)
// }

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
