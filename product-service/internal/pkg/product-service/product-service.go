package productservice

import (
	"fmt"
	"net"
	"product-serivce/internal/pkg/load"
	"product-serivce/internal/service"

	"github.com/Bekzodbekk/zarafshon_city_proto/genproto/productspb"
	"google.golang.org/grpc"
)

type ServiceRun struct {
	serv service.Service
}

func NewServiceRun(serv service.Service) *ServiceRun {
	return &ServiceRun{
		serv: serv,
	}
}

func (s *ServiceRun) RUN(cfg load.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.ProductServiceHost, cfg.Mongosh.Port)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	newServ := grpc.NewServer()
	productspb.RegisterProdcutServiceServer(newServ, &s.serv)

	return newServ.Serve(listener)
}
