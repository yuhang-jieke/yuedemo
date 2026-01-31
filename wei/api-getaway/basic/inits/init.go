package inits

import (
	"flag"
	"log"

	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/config"
	__ "github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/proto"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	ConfigInit()
	GrpcInit()
	RedisInit()
	LogInit()
	middleware.CreateFile()
	EsInit()
}

func GrpcInit() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.UserClient = __.NewUserClient(conn)
}
