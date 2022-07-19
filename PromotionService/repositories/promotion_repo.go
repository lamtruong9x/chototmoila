package repositories

import (
	entity "PromotionService/entities"

	"gorm.io/gorm"
)

type PromotionRepo interface {
	Create(input *entity.Promotion) error
	Get(query string) ([]*entity.Promotion, error)
	Update(input *entity.Promotion) error
	Delete(id string) error
}
type promotionRepo struct {
	db *gorm.DB
}

func NewPromotionRepo(db *gorm.DB) PromotionRepo {
	return &promotionRepo{db: db}
}

func (rp *promotionRepo) Create(input *entity.Promotion) error {
	return rp.db.Create(input).Error
}

func (rp *promotionRepo) Get(query string) ([]*entity.Promotion, error) {
	var results []*entity.Promotion
	err := rp.db.Where(query).Find(&results).Error
	return results, err
}

func (rp *promotionRepo) Update(input *entity.Promotion) error {
	return rp.db.Updates(input).Error
}

func (rp *promotionRepo) Delete(id string) error {
	promotions, err := rp.Get("id = " + id)
	if err != nil {
		return err
	}

	return rp.db.Delete(promotions[0]).Error

}
