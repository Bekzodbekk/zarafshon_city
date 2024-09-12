package service

import (
	"context"
	"product-service/internal/repository"

	"github.com/Bekzodbekk/zarafshon_city_proto/genproto/productspb"
)

type Service struct {
	productspb.UnimplementedProdcutServiceServer
	repo repository.IProductRepository
}

func NewService(repo repository.IProductRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateProduct(ctx context.Context, req *productspb.CreateProductReq) (*productspb.CreateProductResp, error) {
	return nil, nil
}
func (s *Service) UpdateProduct(ctx context.Context, req *productspb.UpdateProductReq) (*productspb.UpdateProductResp, error) {
	return nil, nil
}
func (s *Service) DeleteProduct(ctx context.Context, req *productspb.DeleteProductReq) (*productspb.DeleteProductResp, error) {
	return nil, nil
}
func (s *Service) GetProductById(ctx context.Context, req *productspb.GetProductByIdReq) (*productspb.GetProductByIdResp, error) {
	return nil, nil
}
func (s *Service) GetProducts(ctx context.Context, req *productspb.GetProductsReq) (*productspb.GetProductsResp, error) {
	return nil, nil
}
