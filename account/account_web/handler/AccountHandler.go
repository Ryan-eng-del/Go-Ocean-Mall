package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"net/http"
	"ocean_mall/account/account_srv/proto/pb"
	"ocean_mall/account/account_web/req"
	"ocean_mall/account/account_web/res"
	"ocean_mall/account/internal"
	_jwt "ocean_mall/account/jwt"
	accountLog "ocean_mall/account/log"
	"strconv"
	"time"
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

func LoginByPassword(c *gin.Context) {
	var loginPassword req.LoginByPassword

	err := c.ShouldBind(&loginPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "参数解析错误",
		})
		return
	}
	srvAddr, _ := internal.FilterService("account_srv")
	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure())

	if err != nil {
		accountLog.Logger.Info(fmt.Sprintf("AccountListHandler-Grpc Dial error:%s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := pb.NewAccountServiceClient(conn)

	r, err := client.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: loginPassword.Mobile,
	})

	checkRes, err := client.CheckPassword(context.Background(), &pb.CheckPasswordRequest{
		Password:       loginPassword.Password,
		HashedPassword: r.Password,
		Id:             uint32(r.Id),
	})
	checkResult := "登录失败"

	if checkRes.Result {
		checkResult = "登录成功"

		j := _jwt.NewJWT()

		claims := _jwt.CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "ocean_mall",
				Subject:   "admin",
				Audience:  []string{"somebody_else"},
			},
			AccountId:   r.Id,
			NickName:    r.Nickname,
			AuthorityId: int32(r.Role),
		}

		token, err := j.GenerateJWT(claims)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "generate jwt error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    checkResult,
			"token":  token,
			"result": checkResult,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    checkResult,
		"token":  "",
		"result": checkResult,
	})

}
