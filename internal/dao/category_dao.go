package dao

import (
	"errors"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `json:"nama_kategori" gorm:"not null"`
	Product      *Product
}

type CategoryFilter struct {
	Name          string
	Limit, Offset int
}

func (c *Category) BeforeDelete(tx *gorm.DB) error {
	var count int64
	tx.Model(&Product{}).Where("category_id = ?", c.ID).Count(&count)
	if count > 0 {
		return errors.New("can't delete category with associated product")
	}
	return nil
}
