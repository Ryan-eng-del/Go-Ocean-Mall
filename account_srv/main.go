package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"ocean_mall/account_srv/biz"
	"ocean_mall/account_srv/internal"
	"ocean_mall/account_srv/proto/pb"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入 ip")
	port := flag.Int("port", 9095, "输入端口")
	flag.Parse()
	internal.InitDB()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountService{})
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}

}
