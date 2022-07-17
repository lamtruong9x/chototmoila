package main

import (
	"chotot_product_ltruong/cmd/grpc/protos"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial("0.0.0.0:5000", opts)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// c := protos.NewProductClient(conn)
	// products, _ := GetProduct(c)
}

func GetProduct(c protos.ProductClient) (*protos.GetProductResponse, error) {
	req := protos.GetProductRequest{
		Offset: 0,
	}

	return c.GetProduct(context.Background(), &req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

}
