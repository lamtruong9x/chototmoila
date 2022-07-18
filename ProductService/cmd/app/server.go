package main

import (
	"chotot_product_ltruong/cmd/grpc/protos"
	"chotot_product_ltruong/controller"
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"chotot_product_ltruong/service"
	"chotot_product_ltruong/token"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/mashingan/smapping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	protos.UnimplementedProductServer
	Service service.Service
}

func NewServer(service service.Service) server {
	return server{Service: service}
}

func (s server) Start(port string, jwtManager token.Maker) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	interceptor := controller.NewAuthInterceptor(jwtManager)

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
	}

	server := grpc.NewServer(serverOptions...)
	//server := grpc.NewServer()
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
		//CreatedTime: product.CreatedTime.AsTime(),
		//ExpiredTime: product.ExpiredTime.AsTime(),
		Address: product.Address,
		Content: product.Content,
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

	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "CreateProduct: metadata is not provided")
	}
	//get user id from request
	idString := md.Get(controller.UserIDCtx) 
	if len(idString) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "CreateProduct: metadata is not provided")
	}

	UserID, err := strconv.Atoi(idString[0])
	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateProduct: metadata is not provided")
	}
	productDTO.UserId = UserID

	if err := s.Service.Create(productDTO); err != nil {
		log.Printf("gRPC-CreateProduct: %v", err.Error())
		return nil, status.Error(codes.Internal, s.RichMessage(http.StatusInternalServerError, err))
	}
	return &protos.CreateProductResponse{Message: "created"}, nil
}

// Limit 10 product per call
const (
	Limit = 10
)

func (s server) GetProduct(c context.Context, r *protos.GetProductRequest) (*protos.GetProductResponse, error) {
	offset := int(r.Offset)

	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "GetProduct: metadata is not provided")
	}
	idString := md.Get(controller.UserIDCtx)
	UserID, err := strconv.Atoi(idString[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "GetProduct: metadata is not provided")
	}

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

func (s server) SearchProductBy(c context.Context, r *protos.SearchByRequest) (*protos.SearchByResponse, error) {
	products, err := s.Service.Search(r.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, s.RichMessage(http.StatusNotFound, err))
	}
	protoProducts := make([]*protos.ProductEntity, 0, Limit)
	for _, p := range products {
		protoProducts = append(protoProducts, entityProtoToProto(p))
	}
	return &protos.SearchByResponse{Products: protoProducts}, nil
}
