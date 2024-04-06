package main

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/ishan/Chat_App/db/sqlc"
	"github.com/ishan/Chat_App/pkg/httpserver"
	socket_pkg "github.com/ishan/Chat_App/pkg/websocket"
)

//const connectionStr = "postgres://root:secret@localhost:5432/Realtime_Chat?sslmode=disable"

// const connectionStr = "user=postgres.cztvhelvpqmqvnbriapp password=ishurocks@1502 host=aws-0-ap-south-1.pooler.supabase.com port=5432 dbname=postgres"
const connectionStr = "postgres://root:0d11a0s6MazVFStoPUbR@chat-app.c7gksy8a0o9j.ap-south-1.rds.amazonaws.com:5432/chat_app"

// func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
// 	lc, ok := lambdacontext.FromContext(ctx)
// 	if !ok {
// 		return &events.APIGatewayProxyResponse{
// 			StatusCode: 503,
// 			Body:       "Something went wrong :(",
// 		}, nil
// 	}

// 	cc := lc.ClientContext

// 	return &events.APIGatewayProxyResponse{
// 		StatusCode: 200,
// 		Body:       "Hello, " + cc.Client.AppTitle,
// 	}, nil
// }

func main() {

	conn, err := sql.Open("postgres", connectionStr)

	if err != nil {
		fmt.Println("Error in connecting to Database")
	}

	store := db.NewStore(conn)
	pool := socket_pkg.NewPool()

	httpServer := httpserver.NewHttpServer(store, pool)
	//_ = httpserver.NewHttpServer(store, pool)

	go pool.Start()
	//lambda.Start(httpserver.Handler)

	err1 := httpServer.Run()

	if err1 != nil {
		log.Fatal("Cannot start the server")
	}
}
