// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: product.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Product
	ProductList(ctx context.Context, in *ProductConditionReq, opts ...grpc.CallOption) (*ProductsRes, error)
	BatchGetProduct(ctx context.Context, in *BatchProductIdReq, opts ...grpc.CallOption) (*ProductsRes, error)
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*ProductItemRes, error)
	DeleteProduct(ctx context.Context, in *ProductDelReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProductDetail(ctx context.Context, in *ProductItemReq, opts ...grpc.CallOption) (*ProductItemRes, error)
	// Category
	GetAllCategoriesList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesRes, error)
	GetSubCategory(ctx context.Context, in *CategoriesReq, opts ...grpc.CallOption) (*SubCategoriesRes, error)
	CreateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*CategoryItemRes, error)
	DeleteCategory(ctx context.Context, in *CategoryDelReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Brand
	BrandList(ctx context.Context, in *BrandPagingReq, opts ...grpc.CallOption) (*BrandRes, error)
	CreateBrand(ctx context.Context, in *BrandItemReq, opts ...grpc.CallOption) (*BrandItemRes, error)
	DeleteBrand(ctx context.Context, in *BrandItemReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBrand(ctx context.Context, in *BrandItemReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Advertise
	AdvertiseList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AdvertisesRes, error)
	CreateAdvertise(ctx context.Context, in *AdvertiseReq, opts ...grpc.CallOption) (*AdvertiseItemRes, error)
	DeleteAdvertise(ctx context.Context, in *AdvertiseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateAdvertise(ctx context.Context, in *AdvertiseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// BrandCategory
	CategoryBrandList(ctx context.Context, in *PagingReq, opts ...grpc.CallOption) (*CategoryBrandListRes, error)
	GetCategoryBrandList(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*BrandItemRes, error)
	CreateCategoryBrand(ctx context.Context, in *CategoryBrandReq, opts ...grpc.CallOption) (*CategoryBrandRes, error)
	DeleteCategoryBrand(ctx context.Context, in *CategoryBrandReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCategoryBrand(ctx context.Context, in *CategoryBrandReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) ProductList(ctx context.Context, in *ProductConditionReq, opts ...grpc.CallOption) (*ProductsRes, error) {
	out := new(ProductsRes)
	err := c.cc.Invoke(ctx, "/ProductService/ProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) BatchGetProduct(ctx context.Context, in *BatchProductIdReq, opts ...grpc.CallOption) (*ProductsRes, error) {
	out := new(ProductsRes)
	err := c.cc.Invoke(ctx, "/ProductService/BatchGetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*ProductItemRes, error) {
	out := new(ProductItemRes)
	err := c.cc.Invoke(ctx, "/ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *ProductDelReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductDetail(ctx context.Context, in *ProductItemReq, opts ...grpc.CallOption) (*ProductItemRes, error) {
	out := new(ProductItemRes)
	err := c.cc.Invoke(ctx, "/ProductService/GetProductDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllCategoriesList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesRes, error) {
	out := new(CategoriesRes)
	err := c.cc.Invoke(ctx, "/ProductService/GetAllCategoriesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetSubCategory(ctx context.Context, in *CategoriesReq, opts ...grpc.CallOption) (*SubCategoriesRes, error) {
	out := new(SubCategoriesRes)
	err := c.cc.Invoke(ctx, "/ProductService/GetSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*CategoryItemRes, error) {
	out := new(CategoryItemRes)
	err := c.cc.Invoke(ctx, "/ProductService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteCategory(ctx context.Context, in *CategoryDelReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) BrandList(ctx context.Context, in *BrandPagingReq, opts ...grpc.CallOption) (*BrandRes, error) {
	out := new(BrandRes)
	err := c.cc.Invoke(ctx, "/ProductService/BrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateBrand(ctx context.Context, in *BrandItemReq, opts ...grpc.CallOption) (*BrandItemRes, error) {
	out := new(BrandItemRes)
	err := c.cc.Invoke(ctx, "/ProductService/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteBrand(ctx context.Context, in *BrandItemReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateBrand(ctx context.Context, in *BrandItemReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AdvertiseList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AdvertisesRes, error) {
	out := new(AdvertisesRes)
	err := c.cc.Invoke(ctx, "/ProductService/AdvertiseList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateAdvertise(ctx context.Context, in *AdvertiseReq, opts ...grpc.CallOption) (*AdvertiseItemRes, error) {
	out := new(AdvertiseItemRes)
	err := c.cc.Invoke(ctx, "/ProductService/CreateAdvertise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteAdvertise(ctx context.Context, in *AdvertiseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteAdvertise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateAdvertise(ctx context.Context, in *AdvertiseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/UpdateAdvertise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CategoryBrandList(ctx context.Context, in *PagingReq, opts ...grpc.CallOption) (*CategoryBrandListRes, error) {
	out := new(CategoryBrandListRes)
	err := c.cc.Invoke(ctx, "/ProductService/CategoryBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetCategoryBrandList(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*BrandItemRes, error) {
	out := new(BrandItemRes)
	err := c.cc.Invoke(ctx, "/ProductService/GetCategoryBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateCategoryBrand(ctx context.Context, in *CategoryBrandReq, opts ...grpc.CallOption) (*CategoryBrandRes, error) {
	out := new(CategoryBrandRes)
	err := c.cc.Invoke(ctx, "/ProductService/CreateCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteCategoryBrand(ctx context.Context, in *CategoryBrandReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/DeleteCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateCategoryBrand(ctx context.Context, in *CategoryBrandReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ProductService/UpdateCategoryBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	// Product
	ProductList(context.Context, *ProductConditionReq) (*ProductsRes, error)
	BatchGetProduct(context.Context, *BatchProductIdReq) (*ProductsRes, error)
	CreateProduct(context.Context, *CreateProductReq) (*ProductItemRes, error)
	DeleteProduct(context.Context, *ProductDelReq) (*emptypb.Empty, error)
	UpdateProduct(context.Context, *CreateProductReq) (*emptypb.Empty, error)
	GetProductDetail(context.Context, *ProductItemReq) (*ProductItemRes, error)
	// Category
	GetAllCategoriesList(context.Context, *emptypb.Empty) (*CategoriesRes, error)
	GetSubCategory(context.Context, *CategoriesReq) (*SubCategoriesRes, error)
	CreateCategory(context.Context, *CreateCategoryReq) (*CategoryItemRes, error)
	DeleteCategory(context.Context, *CategoryDelReq) (*emptypb.Empty, error)
	UpdateCategory(context.Context, *CreateCategoryReq) (*emptypb.Empty, error)
	// Brand
	BrandList(context.Context, *BrandPagingReq) (*BrandRes, error)
	CreateBrand(context.Context, *BrandItemReq) (*BrandItemRes, error)
	DeleteBrand(context.Context, *BrandItemReq) (*emptypb.Empty, error)
	UpdateBrand(context.Context, *BrandItemReq) (*emptypb.Empty, error)
	// Advertise
	AdvertiseList(context.Context, *emptypb.Empty) (*AdvertisesRes, error)
	CreateAdvertise(context.Context, *AdvertiseReq) (*AdvertiseItemRes, error)
	DeleteAdvertise(context.Context, *AdvertiseReq) (*emptypb.Empty, error)
	UpdateAdvertise(context.Context, *AdvertiseReq) (*emptypb.Empty, error)
	// BrandCategory
	CategoryBrandList(context.Context, *PagingReq) (*CategoryBrandListRes, error)
	GetCategoryBrandList(context.Context, *CreateCategoryReq) (*BrandItemRes, error)
	CreateCategoryBrand(context.Context, *CategoryBrandReq) (*CategoryBrandRes, error)
	DeleteCategoryBrand(context.Context, *CategoryBrandReq) (*emptypb.Empty, error)
	UpdateCategoryBrand(context.Context, *CategoryBrandReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) ProductList(context.Context, *ProductConditionReq) (*ProductsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (UnimplementedProductServiceServer) BatchGetProduct(context.Context, *BatchProductIdReq) (*ProductsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetProduct not implemented")
}
func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductReq) (*ProductItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *ProductDelReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *CreateProductReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) GetProductDetail(context.Context, *ProductItemReq) (*ProductItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductDetail not implemented")
}
func (UnimplementedProductServiceServer) GetAllCategoriesList(context.Context, *emptypb.Empty) (*CategoriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategoriesList not implemented")
}
func (UnimplementedProductServiceServer) GetSubCategory(context.Context, *CategoriesReq) (*SubCategoriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubCategory not implemented")
}
func (UnimplementedProductServiceServer) CreateCategory(context.Context, *CreateCategoryReq) (*CategoryItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedProductServiceServer) DeleteCategory(context.Context, *CategoryDelReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedProductServiceServer) UpdateCategory(context.Context, *CreateCategoryReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedProductServiceServer) BrandList(context.Context, *BrandPagingReq) (*BrandRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrandList not implemented")
}
func (UnimplementedProductServiceServer) CreateBrand(context.Context, *BrandItemReq) (*BrandItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedProductServiceServer) DeleteBrand(context.Context, *BrandItemReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedProductServiceServer) UpdateBrand(context.Context, *BrandItemReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedProductServiceServer) AdvertiseList(context.Context, *emptypb.Empty) (*AdvertisesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdvertiseList not implemented")
}
func (UnimplementedProductServiceServer) CreateAdvertise(context.Context, *AdvertiseReq) (*AdvertiseItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvertise not implemented")
}
func (UnimplementedProductServiceServer) DeleteAdvertise(context.Context, *AdvertiseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdvertise not implemented")
}
func (UnimplementedProductServiceServer) UpdateAdvertise(context.Context, *AdvertiseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdvertise not implemented")
}
func (UnimplementedProductServiceServer) CategoryBrandList(context.Context, *PagingReq) (*CategoryBrandListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryBrandList not implemented")
}
func (UnimplementedProductServiceServer) GetCategoryBrandList(context.Context, *CreateCategoryReq) (*BrandItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryBrandList not implemented")
}
func (UnimplementedProductServiceServer) CreateCategoryBrand(context.Context, *CategoryBrandReq) (*CategoryBrandRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategoryBrand not implemented")
}
func (UnimplementedProductServiceServer) DeleteCategoryBrand(context.Context, *CategoryBrandReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryBrand not implemented")
}
func (UnimplementedProductServiceServer) UpdateCategoryBrand(context.Context, *CategoryBrandReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategoryBrand not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductConditionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/ProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ProductList(ctx, req.(*ProductConditionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_BatchGetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchProductIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).BatchGetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/BatchGetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).BatchGetProduct(ctx, req.(*BatchProductIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*ProductDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetProductDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductDetail(ctx, req.(*ProductItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllCategoriesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAllCategoriesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetAllCategoriesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAllCategoriesList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetSubCategory(ctx, req.(*CategoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateCategory(ctx, req.(*CreateCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteCategory(ctx, req.(*CategoryDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateCategory(ctx, req.(*CreateCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_BrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandPagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).BrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/BrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).BrandList(ctx, req.(*BrandPagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateBrand(ctx, req.(*BrandItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteBrand(ctx, req.(*BrandItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateBrand(ctx, req.(*BrandItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AdvertiseList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AdvertiseList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/AdvertiseList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AdvertiseList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CreateAdvertise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateAdvertise(ctx, req.(*AdvertiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteAdvertise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteAdvertise(ctx, req.(*AdvertiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/UpdateAdvertise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateAdvertise(ctx, req.(*AdvertiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CategoryBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CategoryBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CategoryBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CategoryBrandList(ctx, req.(*PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetCategoryBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetCategoryBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/GetCategoryBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetCategoryBrandList(ctx, req.(*CreateCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/CreateCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateCategoryBrand(ctx, req.(*CategoryBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/DeleteCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteCategoryBrand(ctx, req.(*CategoryBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateCategoryBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryBrandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateCategoryBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProductService/UpdateCategoryBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateCategoryBrand(ctx, req.(*CategoryBrandReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductList",
			Handler:    _ProductService_ProductList_Handler,
		},
		{
			MethodName: "BatchGetProduct",
			Handler:    _ProductService_BatchGetProduct_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "GetProductDetail",
			Handler:    _ProductService_GetProductDetail_Handler,
		},
		{
			MethodName: "GetAllCategoriesList",
			Handler:    _ProductService_GetAllCategoriesList_Handler,
		},
		{
			MethodName: "GetSubCategory",
			Handler:    _ProductService_GetSubCategory_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _ProductService_CreateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _ProductService_DeleteCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _ProductService_UpdateCategory_Handler,
		},
		{
			MethodName: "BrandList",
			Handler:    _ProductService_BrandList_Handler,
		},
		{
			MethodName: "CreateBrand",
			Handler:    _ProductService_CreateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _ProductService_DeleteBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _ProductService_UpdateBrand_Handler,
		},
		{
			MethodName: "AdvertiseList",
			Handler:    _ProductService_AdvertiseList_Handler,
		},
		{
			MethodName: "CreateAdvertise",
			Handler:    _ProductService_CreateAdvertise_Handler,
		},
		{
			MethodName: "DeleteAdvertise",
			Handler:    _ProductService_DeleteAdvertise_Handler,
		},
		{
			MethodName: "UpdateAdvertise",
			Handler:    _ProductService_UpdateAdvertise_Handler,
		},
		{
			MethodName: "CategoryBrandList",
			Handler:    _ProductService_CategoryBrandList_Handler,
		},
		{
			MethodName: "GetCategoryBrandList",
			Handler:    _ProductService_GetCategoryBrandList_Handler,
		},
		{
			MethodName: "CreateCategoryBrand",
			Handler:    _ProductService_CreateCategoryBrand_Handler,
		},
		{
			MethodName: "DeleteCategoryBrand",
			Handler:    _ProductService_DeleteCategoryBrand_Handler,
		},
		{
			MethodName: "UpdateCategoryBrand",
			Handler:    _ProductService_UpdateCategoryBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
