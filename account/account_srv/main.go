package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"ocean_mall/account/account_srv/biz"
	"ocean_mall/account/account_srv/proto/pb"
	internal2 "ocean_mall/account/internal"
	"ocean_mall/account/util"
	"strconv"
)

func main() {
	port := util.GenRandomPort()
	ip := internal2.ViperConf.AccountSrvConf.Host
	addr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Println(addr, "addr")
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountService{})
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		zap.S().Error("account_srv 启动异常" + err.Error())
		panic(err)
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	internal2.Req(ip, internal2.ViperConf.AccountSrvConf.SrvName, strconv.Itoa(port), port, internal2.ViperConf.AccountSrvConf.Tags, true)
	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}

}
