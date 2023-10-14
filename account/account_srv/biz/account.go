package biz

import (
	"context"
	"crypto/md5"
	"errors"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
	"ocean_mall/account/account_srv/internal"
	"ocean_mall/account/account_srv/model"
	pb2 "ocean_mall/account/account_srv/proto/pb"
	"ocean_mall/account/exception"
)

type AccountService struct {
	pb2.UnimplementedAccountServiceServer
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (a AccountService) GetAccountList(ctx context.Context, request *pb2.PagingRequest) (*pb2.AccountListRes, error) {
	var accountList []model.Account
	result := internal.DB.Scopes(Paginate(int(request.Page), int(request.PageSize))).Find(&accountList)
	if result.Error != nil {
		return nil, result.Error
	}
	accountListRes := &pb2.AccountListRes{}
	accountListRes.Total = int32(result.RowsAffected)
	for _, account := range accountList {
		accountRes := ModelToPb(account)
		accountListRes.AccountList = append(accountListRes.AccountList, accountRes)
	}

	return accountListRes, nil
}

func ModelToPb(account model.Account) *pb2.AccountRes {
	accountRes := &pb2.AccountRes{
		Id:       int32(account.ID),
		Mobile:   account.Mobile,
		Password: account.Password,
		Gender:   account.Gender,
		Role:     uint32(account.Role),
		Nickname: account.NickName,
	}
	return accountRes
}

func (a AccountService) GetAccountByMobile(ctx context.Context, request *pb2.MobileRequest) (*pb2.AccountRes, error) {
	var account model.Account

	result := internal.DB.Where(&model.Account{Mobile: request.Mobile}).First(&account)

	if result.RowsAffected == 0 {
		return nil, errors.New(exception.AccountNotFound)
	}
	res := ModelToPb(account)
	return res, nil
}

func (a AccountService) GetAccountById(ctx context.Context, request *pb2.IdRequest) (*pb2.AccountRes, error) {
	var account model.Account

	result := internal.DB.First(&account, request.Id)

	if result.RowsAffected == 0 {
		return nil, errors.New(exception.AccountNotFound)
	}

	res := ModelToPb(account)
	return res, nil
}

func (a AccountService) AddAccount(ctx context.Context, request *pb2.AddAccountRequest) (*pb2.AccountRes, error) {
	var account model.Account

	result := internal.DB.Where(&model.Account{
		Mobile: request.Mobile,
	}).First(&account)

	if result.RowsAffected != 0 {
		return nil, errors.New(exception.AccountAlreadyExists)
	}

	account.Mobile = request.Mobile
	account.NickName = request.Nickname
	account.Role = 1

	options := password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}

	salt, password := password.Encode(request.Password, &options)
	account.Salt = salt
	account.Password = password
	r := internal.DB.Create(&account)

	if r.Error != nil {
		return nil, errors.New(exception.InternalError)
	}

	accountRes := ModelToPb(account)
	return accountRes, nil
}

func (a AccountService) UpdateAccount(ctx context.Context, request *pb2.UpdateAccountRequest) (*pb2.UpdateAccountRes, error) {
	var account model.Account

	result := internal.DB.First(&account, request.Id)

	if result.RowsAffected == 0 {
		return nil, errors.New(exception.AccountNotFound)
	}

	account.Mobile = request.Mobile
	account.NickName = request.Nickname
	account.Gender = request.Gender

	r := internal.DB.Save(&account)

	if r.Error != nil {
		return nil, errors.New(exception.InternalError)
	}

	return &pb2.UpdateAccountRes{
		Result: true,
	}, nil
}

func (a AccountService) CheckPassword(ctx context.Context, request *pb2.CheckPasswordRequest) (*pb2.CheckPasswordRes, error) {
	var account model.Account

	result := internal.DB.First(&account, request.Id)

	if result.Error != nil {
		return nil, errors.New(exception.InternalError)
	}

	if account.Salt == "" {
		return nil, errors.New(exception.SaltError)
	}

	options := password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}

	r := password.Verify(request.Password, account.Salt, account.Password, &options)

	return &pb2.CheckPasswordRes{
		Result: r,
	}, nil
}
