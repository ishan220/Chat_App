package httpserver

import (
	"fmt"
	_ "math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ishan/Chat_App/db/sqlc"
	"github.com/ishan/Chat_App/model"
	"github.com/lib/pq"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Client   string `json:"client"`
}

type ChatHistoryReq struct {
	FromUser string `json:from_user`
	ToUser   string `json:to_user`
}
type Response struct {
	Status  bool        `json:status`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
}

func (server *HttpServer) AddContact(ctx *gin.Context) {
	fmt.Println("Inside Add Contact Handler")
	ContactToAdd := model.ContactList{}

	err := ctx.ShouldBindJSON(&ContactToAdd)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err = server.store.Queries.VerifyUser(ctx, ContactToAdd.MemberName)

	if err != nil {
		fmt.Println("User not Registered")
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	fmt.Println("Binding to JSON objet successful")
	arg := db.AddContactParams{
		Username:       ContactToAdd.Username,
		MemberUsername: ContactToAdd.MemberName,
		LastActivity:   1211312,
	}
	contactAdded, err := server.store.Queries.AddContact(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, contactAdded)
}

func (server *HttpServer) GetContacts(ctx *gin.Context) {
	username := ctx.Query("username") // shortcut for ctx.Request.URL.Query().Get("lastname")
	contactList, err := server.store.Queries.GetContacts(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, contactList)
}

func (server *HttpServer) ChatBtwnUsers(ctx *gin.Context) {
	from_user := ctx.Query("from")
	to_user := ctx.Query("to")

	if from_user == "" || to_user == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "from_user or to_user is empty"})
		return
	}

	arg := db.GetChatHistoryParams{
		FromUser: from_user,
		ToUser:   to_user,
	}

	chats, err := server.store.GetChatHistory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, chats)

}

func (server *HttpServer) CreateChat(ctx *gin.Context) {
	fmt.Println("Inside http server createChat Handler")
	chat, _ := ctx.Get("chat")
	fmt.Printf("After Must Get %+v", chat)
	arg := db.CreateChatParams{
		//ID:        strconv.FormatInt(RandomInt(0, 100000), 10),
		ID:        chat.(*model.Chat).ID,
		FromUser:  chat.(*model.Chat).FromUser,
		ToUser:    chat.(*model.Chat).ToUser,
		Message:   chat.(*model.Chat).Message,
		CreatedAt: time.Now(),
	}

	chatCreated, err := server.store.Queries.CreateChat(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	fmt.Println("Chat Created", chatCreated)

	ctx.JSON(http.StatusOK, chatCreated)

}

func (server *HttpServer) loginUser(ctx *gin.Context) {
	fmt.Println("Inside http server loginUser Handler")
	userReq := &UserReq{}
	err := ctx.ShouldBindJSON(userReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}
	arg := db.GetUserParams{
		Username: userReq.Username,
		Password: userReq.Password,
	}
	userLoggedIn, err := server.store.Queries.GetUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, userLoggedIn)
}

func (server *HttpServer) createUser(ctx *gin.Context) {
	fmt.Println("Inside http server createUser Handler")
	userReq := &UserReq{}
	err := ctx.ShouldBindJSON(userReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}
	arg := db.CreateUserParams{
		Username: userReq.Username,
		Password: userReq.Password,
	}
	userCreated, err := server.store.Queries.CreateUser(ctx, arg)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, userCreated)
}
