package init

import (
	"flag"
	"log"
	"zholianxi/bff/basic/config"
	"zholianxi/bff/middleware"
	__ "zholianxi/src/basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	InitViper()
	InitMysql()
	InitRedis()
	InitEs()
	LogInit()
	initDB()
	middleware.CreateFile()
}

func initDB() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	config.GoodsClient = __.NewGoodsClient(conn)
}
