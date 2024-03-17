docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
migrate create -ext sql -dir db/migration/ -seq add_users
migrate -path=db/migration -database "postgres://root:secret@localhost:5432/Realtime_Chat?sslmode=disable" -verbose up
gmake sqlc(sqlc generate)
go get github.com/lib/pq  //for sql postgres driver