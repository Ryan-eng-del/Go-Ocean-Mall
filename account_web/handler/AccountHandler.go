package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"ocean_mall/account_srv/proto/pb"
	"ocean_mall/account_web/res"
	accountLog "ocean_mall/log"
	"strconv"
)

func AccountListHandler(c *gin.Context) {
	accountLog.Logger.Info("---- AccountListHandler -----")
	page := c.DefaultQuery("page", "1")
	pageNum := c.DefaultQuery("page_size", "10")

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())

	if err != nil {
		accountLog.Logger.Info(fmt.Sprintf("AccountListHandler-Grpc Dial error:%s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

	client := pb.NewAccountServiceClient(conn)
	pageNo, _ := strconv.ParseInt(page, 10, 32)
	pageSize, _ := strconv.ParseInt(pageNum, 10, 32)
	fmt.Println(pageNo, pageSize)
	list, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageSize: uint32(pageSize),
		Page:     uint32(pageNo),
	})

	if err != nil {
		accountLog.Logger.Info(fmt.Sprintf("AccountListHandler 调用失败: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

	var AccountResList []res.AccountRes

	for _, account := range list.AccountList {
		AccountResList = append(AccountResList, pbToRes(account))
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "Success",
		"data":  AccountResList,
		"total": list.Total})
}

func pbToRes(accountRes *pb.AccountRes) res.AccountRes {
	return res.AccountRes{
		Mobile:   accountRes.Mobile,
		NickName: accountRes.Nickname,
		Gender:   accountRes.Gender,
	}

}
