package main

import (
	"chotot_product_ltruong/cmd/grpc/protos"
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"chotot_product_ltruong/service"
	"context"
	"fmt"
	"github.com/mashingan/smapping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s server) Start(port string) {
	server := grpc.NewServer()
	protos.RegisterProductServer(server, s)
	reflection.Register(server)
	address := fmt.Sprintf("0.0.0.0:%s", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(server.Serve(lis))
}

// c for grpc error code and h is for http status code
func (s server) RichMessage(c int, err error) string {
	return fmt.Sprintf("%v;%v", c, err.Error())
}

func productProtoToEntity(product *protos.CreateProductRequest) *entity.Product {
	return &entity.Product{
		ProductName: product.ProductName,
		UserId:      1,
		CatId:       product.CatId,
		TypeId:      product.TypeId,
		Price:       product.Price,
		State:       product.State,
		CreatedTime: product.CreatedTime.AsTime(),
		ExpiredTime: product.ExpiredTime.AsTime(),
		Address:     product.Address,
		Content:     product.Content,
	}
}

func entityProtoToProto(product *entity.Product) *protos.ProductEntity {
	return &protos.ProductEntity{
		Id:          int32(product.Id),
		ProductName: product.ProductName,
		CatId:       product.CatId,
		TypeId:      product.TypeId,
		Price:       product.Price,
		State:       product.State,
		CreatedTime: timestamppb.New(product.CreatedTime),
		ExpiredTime: timestamppb.New(product.ExpiredTime),
		Priority:    product.Priority,
		Address:     product.Address,
		Content:     product.Content,
		Images:      nil,
	}
}

func (s server) CreateProduct(c context.Context, r *protos.CreateProductRequest) (*protos.CreateProductResponse, error) {

	product := productProtoToEntity(r)
	productDTO := &dto.Product{}
	if err := smapping.FillStruct(productDTO, smapping.MapFields(product)); err != nil {
		log.Printf("gRPC-CreateProduct-smapping: %v", err.Error())
		return nil, status.Error(codes.Internal, s.RichMessage(http.StatusInternalServerError, err))
	}

	if err := s.Service.Create(productDTO); err != nil {
		log.Printf("gRPC-CreateProduct: %v", err.Error())
		return nil, status.Error(codes.Internal, s.RichMessage(http.StatusInternalServerError, err))
	}
	return &protos.CreateProductResponse{Message: "created"}, nil
}

// Limit 10 product per call
const (
	Limit  = 10
	UserID = 1
)

func (s server) GetProduct(c context.Context, r *protos.GetProductRequest) (*protos.GetProductResponse, error) {
	offset := int(r.Offset)
	products, err := s.Service.GetByUserID(UserID, Limit, offset)
	if err != nil {
		log.Printf("gRPC-GetProduct: %v", err.Error())
		return nil, status.Error(codes.Internal, s.RichMessage(http.StatusInternalServerError, err))
	}

	protoProducts := make([]*protos.ProductEntity, 0, Limit)
	for _, p := range products {
		protoProducts = append(protoProducts, entityProtoToProto(p))
	}

	return &protos.GetProductResponse{Products: protoProducts}, nil
}

func (s server) UpdateProduct(c context.Context, r *protos.UpdateProductRequest) (*protos.ProductEntity, error) {
	return nil, nil
}
