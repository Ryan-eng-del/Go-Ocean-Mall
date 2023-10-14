package biz

import (
	"context"
	"fmt"
	"ocean_mall/account/account_srv/proto/pb"
	"ocean_mall/account/internal"
	"testing"
)

func init() {
	internal.InitDB()
}

func TestAccountServer_AddAccount(t *testing.T) {
	accountServer := AccountService{}

	for i := 0; i < 5; i++ {
		s := fmt.Sprintf("13000000000%d", i)

		res, err := accountServer.AddAccount(context.Background(), &pb.AddAccountRequest{
			Mobile:   s,
			Password: s,
			Nickname: s,
			Gender:   "male",
		})

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res.Id)
	}

}

func TestAccountServer_GetAccountList(t *testing.T) {
	accountServer := AccountService{}

	res, err := accountServer.GetAccountList(context.Background(), &pb.PagingRequest{
		Page:     1,
		PageSize: 3,
	})

	if err != nil {
		fmt.Println(err)
	}
	for _, account := range res.AccountList {
		fmt.Println(account.Id)
	}

	res, err = accountServer.GetAccountList(context.Background(), &pb.PagingRequest{
		Page:     2,
		PageSize: 3,
	})

	if err != nil {
		fmt.Println(err)
	}
	for _, account := range res.AccountList {
		fmt.Println(account.Id)
	}
}

func TestAccountServer_GetAccountByMobile(t *testing.T) {
	mobile := "130000000003"

	accountServer := AccountService{}
	res, err := accountServer.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: mobile,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Id)
}

func TestAccountServer_GetAccountById(t *testing.T) {
	id := 1

	accountService := AccountService{}

	res, err := accountService.GetAccountById(context.Background(), &pb.IdRequest{
		Id: uint32(id),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Id)
}
