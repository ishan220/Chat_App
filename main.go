package main

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/ishan/Chat_App/db/sqlc"
	"github.com/ishan/Chat_App/pkg/httpserver"
	socket_pkg "github.com/ishan/Chat_App/pkg/websocket"
)

const connectionStr = "postgres://root:secret@localhost:5432/Realtime_Chat?sslmode=disable"

// func init() {

// }
func main() {

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("Error in connecting to Database")
	}
	store := db.NewStore(conn)
	pool := socket_pkg.NewPool()

	httpServer := httpserver.NewHttpServer(store, pool)
	go pool.Start()
	err1 := httpServer.Run()
	if err1 != nil {
		log.Fatal("Cannot start the server")
	}
}
