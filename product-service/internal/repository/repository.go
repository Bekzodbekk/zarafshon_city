package repository

import (
	"context"

	"github.com/Bekzodbekk/zarafshon_city_proto/genproto/productspb"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, req *productspb.CreateProductReq) (*productspb.CreateProductResp, error)
	UpdateProduct(ctx context.Context, req *productspb.UpdateProductReq) (*productspb.UpdateProductResp, error)
	DeleteProduct(ctx context.Context, req *productspb.DeleteProductReq) (*productspb.DeleteProductResp, error)
	GetProductById(ctx context.Context, req *productspb.GetProductByIdReq) (*productspb.GetProductByIdResp, error)
	GetProducts(ctx context.Context, req *productspb.GetProductsReq) (*productspb.GetProductsResp, error)
}
