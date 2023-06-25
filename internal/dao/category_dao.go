package dao

import (
	"errors"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `json:"nama_kategori"`
	Product      *Product
}

func (c *Category) BeforeDelete(tx *gorm.DB) error {
	var count int64
	tx.Model(&Product{}).Where("category_id = ?", c.ID).Count(&count)
	if count > 0 {
		return errors.New("can't delete category with associated product")
	}
	return nil
}
