package main

import (
	"chotot_product_ltruong/config"
	"chotot_product_ltruong/controller"
	"chotot_product_ltruong/repo"
	"chotot_product_ltruong/service"
	"chotot_product_ltruong/token"
	"log"
)

type application struct {
	Controller controller.Controller
}

const (
	PORT      = ":8081"
	GRPC_PORT = "5000"
	secretKey = "3y6B8DbGdJfNjQmSqVsXu2x4z7C9EbHeKgNr"
)

func main() {

	// Open DB
	db := config.InitDB()
	defer config.CloseDB(db)

	// Create repo
	repo1 := repo.New(db)
	svc := service.New(&repo1)
	// create jwt maker
	maker, err := token.NewJWTMaker(secretKey)
	//log.Fatal(err)
	ctrl := controller.New(svc, maker)
	app := application{
		Controller: ctrl,
	}
	// grpc
	grpc := NewServer(svc)
	go grpc.Start(GRPC_PORT, maker)
	r := app.NewRouter()

	err = r.Run(PORT)
	log.Fatal(err)

}
