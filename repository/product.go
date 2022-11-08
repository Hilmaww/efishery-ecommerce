package repository

import (
	"efishery-ecommerce/entity"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetProductsById(id string) (entity.Product, error)
	GetFiltered(category string, hargaterendah string, hargatertinggi string) ([]entity.Product, error)
	AddToCart(entity.Cart) error
	GetCartAll() ([]entity.Cart, error)
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

func (u ProductRepository) GetFiltered(category string, hargaterendah string, hargatertinggi string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if hargaterendah == "" {
		hargaterendah = "0"
	}
	if hargatertinggi == "" {
		hargatertinggi = "9999999999999"
	}

	if category == "" {
		err = u.db.Debug().Where("Harga <= ? AND Harga >= ?", hargatertinggi, hargaterendah).Find(&products).Error
	} else {
		err = u.db.Debug().Where("Kategori = ? AND Harga <= ? AND Harga >= ?", category, hargatertinggi, hargaterendah).Find(&products).Error
	}

	if err != nil {
		//fmt.Println("hayo gagal filter")
		return nil, err
	}
	return products, nil
}

func (u ProductRepository) AddToCart(cart entity.Cart) error {
	if err := u.db.Create(&cart).Error; err != nil {
		return err
	}
	return nil
}
func (u ProductRepository) GetCartAll() ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := u.db.Debug().Find(&carts).Error; err != nil {
		return nil, nil
	}
	return carts, nil
}
