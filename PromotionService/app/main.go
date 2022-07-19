package main

import (
	"PromotionService/config"
	"PromotionService/controller"
	"PromotionService/repositories"
	"PromotionService/services"
	"PromotionService/token"
	"log"
)

type application struct {
	Controller controller.PromotionController
}

const (
	PORT      = ":8082"
	secretKey = "3y6B8DbGdJfNjQmSqVsXu2x4z7C9EbHeKgNr"
)

var (
	// Open DB
	db = config.InitDB()

	//create repo, svc, ctrl
	repo     = repositories.NewPromotionRepo(db)
	svc      = services.NewPromotionService(repo)
	maker, _ = token.NewJWTMaker(secretKey)
	ctrl     = controller.NewPromotionController(svc, maker)

	app = application{
		Controller: ctrl,
	}
)

func main() {
	defer config.CloseDB(db)
	r := app.NewRouter()

	err := r.Run(PORT)
	log.Fatal(err)
}
