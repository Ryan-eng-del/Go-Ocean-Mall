package biz

import (
	"context"
	"ocean_mall/account_srv/proto/pb"
)

type AccountService struct {
}

func (a AccountService) GetAccountList(ctx context.Context, request *pb.PagingRequest) (*pb.AccountListRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) GetAccountByMobile(ctx context.Context, request *pb.MobileRequest) (*pb.AccountRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) GetAccountById(ctx context.Context, request *pb.IdRequest) (*pb.AccountRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) AddAccount(ctx context.Context, request *pb.AddAccountRequest) (*pb.AccountRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) UpdateAccount(ctx context.Context, request *pb.UpdateAccountRequest) (*pb.UpdateAccountRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) CheckPassword(ctx context.Context, request *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountService) mustEmbedUnimplementedAccountServiceServer() {
	//TODO implement me
	panic("implement me")
}
