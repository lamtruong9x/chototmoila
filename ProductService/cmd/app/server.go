package main

import (
	"chotot_product_ltruong/cmd/grpc/protos"
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/service"
	"context"
	"fmt"
	"github.com/mashingan/smapping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
)

type server struct {
	protos.UnimplementedProductServer
	Service service.Service
}

func NewServer(service service.Service) server {
	return server{Service: service}
}

func (s server) Start(port string) error {
	server := grpc.NewServer()
	protos.RegisterProductServer(server, s)
	reflection.Register(server)
	address := fmt.Sprintf("0.0.0.0:%s", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return server.Serve(lis)
}

// c for grpc error code and h is for http status code
func (s server) RichMessage(c int, err error) string {
	return fmt.Sprintf("%v;%v", c, err.Error())
}
func (s server) CreateProduct(c context.Context, r *protos.CreateProductRequest) (*protos.CreateProductResponse, error) {
	var product *dto.Product
	if err := smapping.FillStruct(product, smapping.MapFields(r)); err != nil {
		log.Printf("gRPC-CreateProduct: %v", err.Error())
		return nil, status.Error(codes.Internal, s.RichMessage(http.StatusInternalServerError, err))
	}
	if err := s.Service.Create(product); err != nil {
		log.Printf("gRPC-CreateProduct: %v", err.Error())
		return nil, status.Error(codes.Internal, s.RichMessage(http.StatusInternalServerError, err))
	}
	return &protos.CreateProductResponse{Message: "created"}, nil
}
