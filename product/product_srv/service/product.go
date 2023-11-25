package service

import (
	"context"
	"ocean_mall/product/product_srv/proto/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)
type Product struct {
	pb.UnimplementedProductServiceServer
}

// Product
func (mt Product) ProductList(_ context.Context, _ *pb.ProductConditionReq) (*pb.ProductsRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) BatchGetProduct(_ context.Context, _ *pb.BatchProductIdReq) (*pb.ProductsRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) CreateProduct(_ context.Context, _ *pb.CreateProductReq) (*pb.ProductItemRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) DeleteProduct(_ context.Context, _ *pb.ProductDelReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) UpdateProduct(_ context.Context, _ *pb.CreateProductReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) GetProductDetail(_ context.Context, _ *pb.ProductItemReq) (*pb.ProductItemRes, error) {
	panic("not implemented") // TODO: Implement
}

// Category
func (mt Product) GetAllCategoriesList(_ context.Context, _ *emptypb.Empty) (*pb.CategoriesRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) GetSubCategory(_ context.Context, _ *pb.CategoriesReq) (*pb.SubCategoriesRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) CreateCategory(_ context.Context, _ *pb.CreateCategoryReq) (*pb.CategoryItemRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) DeleteCategory(_ context.Context, _ *pb.CategoryDelReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) UpdateCategory(_ context.Context, _ *pb.CreateCategoryReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

// Brand
func (mt Product) BrandList(_ context.Context, _ *pb.BrandPagingReq) (*pb.BrandRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) CreateBrand(_ context.Context, _ *pb.BrandItemReq) (*pb.BrandItemRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) DeleteBrand(_ context.Context, _ *pb.BrandItemReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) UpdateBrand(_ context.Context, _ *pb.BrandItemReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

// Advertise
func (mt Product) AdvertiseList(_ context.Context, _ *emptypb.Empty) (*pb.AdvertisesRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) CreateAdvertise(_ context.Context, _ *pb.AdvertiseReq) (*pb.AdvertiseItemRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) DeleteAdvertise(_ context.Context, _ *pb.AdvertiseReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) UpdateAdvertise(_ context.Context, _ *pb.AdvertiseReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

// BrandCategory
func (mt Product) CategoryBrandList(_ context.Context, _ *pb.PagingReq) (*pb.CategoryBrandListRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) GetCategoryBrandList(_ context.Context, _ *pb.CreateCategoryReq) (*pb.BrandItemRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) CreateCategoryBrand(_ context.Context, _ *pb.CategoryBrandReq) (*pb.CategoryBrandRes, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) DeleteCategoryBrand(_ context.Context, _ *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) UpdateCategoryBrand(_ context.Context, _ *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (mt Product) mustEmbedUnimplementedProductServiceServer() {
	panic("not implemented") // TODO: Implement
}

