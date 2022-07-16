package main

import (
	"chotot_product_ltruong/config"
	"chotot_product_ltruong/controller"
	"chotot_product_ltruong/repo"
	"chotot_product_ltruong/service"
	"log"
)

type applicatiton struct {
	Controller controller.Controller
}

const (
	PORT      = ":8080"
	GRPC_PORT = "5000"
)

func main() {

	// Open DB
	db := config.InitDB()
	defer config.CloseDB(db)

	// Create repo
	repo1 := repo.New(db)
	svc := service.New(&repo1)
	ctrl := controller.New(svc)
	app := applicatiton{
		Controller: ctrl,
	}
	// grpc
	//db1 := config.InitDB()
	//defer config.CloseDB(db1)
	//repo2 := repo.New(db1)
	//svc1 := service.New(&repo2)
	grpc := NewServer(svc)
	go grpc.Start(GRPC_PORT)
	r := app.NewRouter()

	err := r.Run(PORT)
	log.Fatal(err)

}
