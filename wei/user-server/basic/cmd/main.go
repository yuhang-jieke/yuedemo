package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/yuhang-jieke/yuedemo/wei/user-server/basic/init"
	__ "github.com/yuhang-jieke/yuedemo/wei/user-server/handler/proto"
	"github.com/yuhang-jieke/yuedemo/wei/user-server/handler/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

// server is used to implement helloworld.GreeterServer.

func main() {
	//config.RDB.Set(context.Background(), "k", "1", 1)
	//str := config.RDB.Get(context.Background(), "k")
	/*config.RDB.Set(context.Background(), "k1", "1111", 10*time.Second)
	str := config.RDB.Get(context.Background(), "k1")
	fmt.Println("------------------->", str.String())*/

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterUserServer(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
