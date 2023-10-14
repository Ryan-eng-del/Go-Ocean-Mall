package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ocean_mall/account/account_web/handler"
	"ocean_mall/account/account_web/middleware"
	"ocean_mall/account/internal"
	"strconv"
)

func main() {
	ip := flag.String("ip", internal.ViperConf.AccountWebConf.Host, "输入 ip")
	port := flag.String("port", internal.ViperConf.AccountWebConf.Port, "输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	fmt.Println(addr, "addr")
	r := gin.Default()
	r.Use(middleware.Cors())
	p, _ := strconv.ParseInt(*port, 10, 32)
	accountGroup := r.Group("/v1/account")
	{
		accountGroup.GET("/list", handler.AccountListHandler)
		accountGroup.POST("/login", handler.LoginByPassword)

	}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
	internal.Req(*ip, internal.ViperConf.AccountWebConf.SrvName, *port, int(p), internal.ViperConf.AccountWebConf.Tags, false)
	r.Run(addr)

}
