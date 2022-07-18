package entity

import "time"

type PromotionEntity struct {
	Id          int       `json:"id" gorm:"primary_key;auto_increment"`
	ProductId   int       `json:"product_id" gorm:"type:int"`
	CreatedTime time.Time `gorm:"type:timestamp,autoCreateTime"`
	ExpiredTime time.Time `gorm:"type:timestamp"`
	LastUpdate  time.Time `gorm:"type:timestamp"`
}
