package repository

import (
	"efishery-ecommerce/entity"
	"fmt"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetProductsById(id string) (entity.Product, error)
	AddToCart([]entity.Cart) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (u ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := u.db.Debug().Find(&products).Error; err != nil {
		return nil, nil
	}
	return products, nil
}
func (u ProductRepository) GetProductsById(id string) (entity.Product, error) {
	var products entity.Product
	if err := u.db.Debug().First(&products, id).Error; err != nil {
		return products, nil
	}
	return products, nil
}

func (u ProductRepository) AddToCart(cart []entity.Cart) error {
	if err := u.db.Create(&cart).Error; err != nil {
		fmt.Println("gagal disini")
		return err
	}
	return nil
}
