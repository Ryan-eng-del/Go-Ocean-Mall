package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"ocean_mall/account_web/handler"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入 ip")
	port := flag.Int("port", 8085, "输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	fmt.Println(addr, "addr")
	r := gin.Default()

	accountGroup := r.Group("/v1/account")

	{
		accountGroup.GET("/list", handler.AccountListHandler)
	}

	r.Run(addr)
}
