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

func (r *repo) GetByID(productID int) (*entity.Product, error) {
	var product entity.Product
	tx := r.DB.First(&product, productID)
	if err := tx.Error; err != nil {
		fmt.Printf("buggggg")
		return nil, err
	}
	return &product, nil
}

func (r *repo) Update(userID int, product *entity.Product) (*entity.Product, error) {
	var output entity.Product

	// Check if the owner of the product is valid one
	if userID != product.UserId {
		return nil, errors.New("invalid user")
	}

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

	product, err = r.GetByID(product.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *repo) Delete(productID, userID int) error {
	//if err := r.DB.Delete(&entity.Product{}, id).Error; err != nil {
	//	return err
	//}
	if err := r.DB.Where("id = ? AND user_id = ?", productID, userID).Delete(&entity.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) Search(query string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := r.DB.Where(query).Find(&products).Error
	return products, err
}
