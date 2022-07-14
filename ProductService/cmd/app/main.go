package main

import (
	"chotot_product_ltruong/config"
	"chotot_product_ltruong/controller"
	repo2 "chotot_product_ltruong/repo"
	"chotot_product_ltruong/service"
	"log"
)

type applicatiton struct {
	Controller controller.Controller
}

const PORT = ":8080"

func main() {

	// Open DB
	db := config.InitDB()
	defer config.CloseDB(db)

	// Create repo
	repo := repo2.New(db)
	svc := service.New(&repo)
	ctrl := controller.New(svc)
	app := applicatiton{
		Controller: ctrl,
	}
	r := app.NewRouter()
	err := r.Run(PORT)
	log.Fatal(err)
}
