package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"ocean_mall/account/account_srv/biz"
	"ocean_mall/account/account_srv/proto/pb"
	internal2 "ocean_mall/account/internal"
	"strconv"
)

type AccountSrvConf struct {
}

func main() {
	ip := flag.String("ip", internal2.ViperConf.AccountSrvConf.Host, "输入 ip")
	port := flag.String("port", internal2.ViperConf.AccountSrvConf.Port, "输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	fmt.Println(addr, "addr")

	internal2.InitDB()
	internal2.InitRedis()
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountService{})
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		zap.S().Error("account_srv 启动异常" + err.Error())
		panic(err)
	}
	p, _ := strconv.ParseInt(*port, 10, 32)
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	internal2.Req(*ip, internal2.ViperConf.AccountSrvConf.SrvName, *port, int(p), internal2.ViperConf.AccountSrvConf.Tags, true)
	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}

}
