package service

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
)

type Service interface {
	Create(product *dto.Product) error
	GetByID(productID int) (*entity.Product, error)
	GetByUserID(userID int, limit, offset int) ([]*entity.Product, error)
	Update(product *dto.ProductUpdate) (*entity.Product, error)
	Delete(id int) error
	Seach(query string) ([]*entity.Product, error)
}
