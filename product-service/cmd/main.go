package main

import (
	"log"
	"product-service/internal/pkg/load"
	"product-service/internal/pkg/mongosh"
	productservice "product-service/internal/pkg/product-service"
	"product-service/internal/repository"
	"product-service/internal/service"
)

func main() {
	cfg, err := load.Load("./config")
	if err != nil {
		log.Fatal(err)
	}

	mongoshStruct, err := mongosh.ConnectMongo(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewProductRepo(mongoshStruct.Collection)
	service := service.NewService(repo)
	runService := productservice.NewServiceRun(*service)

	log.Printf("Product Service Running on :%d port", cfg.ProductServicePort)
	if err := runService.RUN(*cfg); err != nil {
		log.Fatal(err)
	}
}
