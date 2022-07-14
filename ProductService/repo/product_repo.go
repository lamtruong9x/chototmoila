package repo

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"errors"
	"fmt"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repo {
	return repo{
		DB: db,
	}
}

func (r *repo) Create(product *dto.Product) error {
	if err := r.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) GetByUserID(userID int, limit, offset int) ([]*entity.Product, error) {
	var products []*entity.Product
	tx := r.DB.Where("user_id=?", userID).Limit(limit).Offset(offset).Find(&products).Order("created_time")
	if err := tx.Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repo) Update(userID int, product *entity.Product) (*entity.Product, error) {
	var output entity.Product

	// Check if the owner of the product is valid one
	if userID != product.UserId {
		return nil, errors.New("invalid user")
	}
	//fmt.Println(product.Id)
	if err := r.DB.Where("id = ?", product.Id).First(&output).Error; err != nil {
		return nil, err
	}
	err := smapping.FillStruct(&output, smapping.MapFields(product))
	if err != nil {
		return nil, err
	}
	if err := r.DB.Model(&output).Updates(&product).Error; err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", output)
	return &output, nil
}

func (r *repo) Delete(id int) error {
	if err := r.DB.Delete(&entity.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
