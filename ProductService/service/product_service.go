package service

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"chotot_product_ltruong/repo"
	"time"

	"github.com/mashingan/smapping"
)

type service struct {
	Repo repo.Repo
}

func New(repo repo.Repo) *service {
	return &service{Repo: repo}
}

const EXPIRED_TIME = 24 * 30

func (svc *service) Create(product *dto.Product) error {
	product.CreatedTime = time.Now()
	product.ExpiredTime = time.Now().Add(time.Hour * EXPIRED_TIME)
	return svc.Repo.Create(product)
}

func (svc *service) GetByID(productID int) (*entity.Product, error) {
	return svc.Repo.GetByID(productID)
}

func (svc *service) GetByUserID(userID int, limit, offset int) ([]*entity.Product, error) {
	offset *= limit
	return svc.Repo.GetByUserID(userID, limit, offset)
}

func (svc *service) Update(product *dto.ProductUpdate) (*entity.Product, error) {
	input := &entity.Product{}
	err := smapping.FillStruct(input, smapping.MapFields(product))
	if err != nil {
		return nil, err
	}
	return svc.Repo.Update(1, input)
}

func (svc *service) Delete(id int) error {
	return svc.Repo.Delete(id)
}

func (svc *service) Search(query string) ([]*entity.Product, error) {
	return svc.Repo.Search(query)
}
