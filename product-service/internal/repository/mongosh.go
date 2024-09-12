package repository

import (
	"context"

	"github.com/Bekzodbekk/zarafshon_city_proto/genproto/productspb"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	collection *mongo.Collection
}

func NewProductRepo(coll *mongo.Collection) IProductRepository {
	return &ProductRepo{
		collection: coll,
	}
}

func (p *ProductRepo) CreateProduct(ctx context.Context, req *productspb.CreateProductReq) (*productspb.CreateProductResp, error) {
	return nil, nil
}
func (p *ProductRepo) UpdateProduct(ctx context.Context, req *productspb.UpdateProductReq) (*productspb.UpdateProductResp, error) {
	return nil, nil
}
func (p *ProductRepo) DeleteProduct(ctx context.Context, req *productspb.DeleteProductReq) (*productspb.DeleteProductResp, error) {
	return nil, nil
}
func (p *ProductRepo) GetProductById(ctx context.Context, req *productspb.GetProductByIdReq) (*productspb.GetProductByIdResp, error) {
	return nil, nil
}
func (p *ProductRepo) GetProducts(ctx context.Context, req *productspb.GetProductsReq) (*productspb.GetProductsResp, error) {
	return nil, nil
}
