package repo

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
)

type Repo interface {
	Create(product *dto.Product) error
	GetByUserID(userID int, limit, offset int) ([]*entity.Product, error)
	Update(userID int, product *entity.Product) (*entity.Product, error)
	Delete(id int) error
	GetByID(productID int) (*entity.Product, error)
}
