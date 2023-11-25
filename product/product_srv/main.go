package main

import (
	"fmt"
	"net"
	"ocean_mall/product/internal"
	"ocean_mall/product/product_srv/proto/pb"
	"ocean_mall/product/product_srv/service"
	"ocean_mall/product/util"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	port := util.GenRandomPort()
	ip := internal.ViperConf.ProductSrvConf.Host
	addr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Println(addr, "addr")
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, &service.Product{})
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		zap.S().Error("account_srv 启动异常" + err.Error())
		panic(err)
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	internal.Req(ip, internal.ViperConf.ProductSrvConf.SrvName, strconv.Itoa(port), port, internal.ViperConf.ProductSrvConf.Tags, true)
	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}

}
