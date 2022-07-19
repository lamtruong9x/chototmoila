package services

import (
	entity "PromotionService/entities"
	"PromotionService/repositories"
)

type PromotionService interface {
	Create(input *entity.Promotion) error
	Get(query string) ([]*entity.Promotion, error)
	Update(input *entity.Promotion) error
	Delete(id string) error
}

type promotionService struct {
	repo repositories.PromotionRepo
}

func NewPromotionService(repo repositories.PromotionRepo) PromotionService {
	return &promotionService{repo: repo}
}

func (svc *promotionService) Create(input *entity.Promotion) error {
	return svc.repo.Create(input)
}

func (svc *promotionService) Get(query string) ([]*entity.Promotion, error) {
	return svc.repo.Get(query)
}

func (svc *promotionService) Update(input *entity.Promotion) error {
	return svc.repo.Update(input)
}

func (svc *promotionService) Delete(id string) error {
	return svc.repo.Delete(id)
}
