package service

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
)

type Service interface {
	Create(product *dto.Product) error
	//Get(any)
	GetByUserID(userID int, limit, offset int) ([]*entity.Product, error)
	Update(product *dto.ProductUpdate) (*entity.Product, error)
	Delete(id int) error
}
