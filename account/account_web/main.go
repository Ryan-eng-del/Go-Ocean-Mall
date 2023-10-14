package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"ocean_mall/account/account_web/handler"
	"ocean_mall/account/internal"
)

func main() {
	ip := flag.String("ip", internal.ViperConf.AccountWebConf.Host, "输入 ip")
	port := flag.String("port", internal.ViperConf.AccountWebConf.Port, "输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	fmt.Println(addr, "addr")
	r := gin.Default()

	accountGroup := r.Group("/v1/account")

	{
		accountGroup.GET("/list", handler.AccountListHandler)
		accountGroup.POST("/login", handler.LoginByPassword)
	}

	r.Run(addr)
}
