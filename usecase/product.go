package usecase

import (
	"efishery-ecommerce/entity"
	"efishery-ecommerce/entity/response"
	"efishery-ecommerce/repository"
	"github.com/jinzhu/copier"
)

type IProductUsecase interface {
	GetProductList() ([]response.GetProductResponse, error)
	GetProductDetail(id string) ([]response.GetProductDetailResponse, error)
	GetFilteredProduct(category string, hargaterendah string, hargatertinggi string) ([]response.GetProductResponse, error)
	AddProductToCart(req *response.AddToCartRequest) (response.AddToCartRequest, error)
	GetCartList() ([]response.GetCartResponse, error)
}
type ProductUsecase struct {
	productRepository repository.IProductRepository
}

func NewProductUsecase(productRepository repository.IProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepository}
}

func (u ProductUsecase) GetProductList() ([]response.GetProductResponse, error) {
	products, err := u.productRepository.GetAll()
	if err != nil {
		return nil, nil
	}
	var productResponse []response.GetProductResponse
	copier.Copy(&productResponse, &products)
	return productResponse, nil
}

func (u ProductUsecase) GetProductDetail(id string) ([]response.GetProductDetailResponse, error) {
	products, err := u.productRepository.GetProductsById(id)
	if err != nil {
		return nil, nil
	}
	var productResponse []response.GetProductDetailResponse
	copier.Copy(&productResponse, &products)
	return productResponse, nil
}

func (u ProductUsecase) AddProductToCart(req *response.AddToCartRequest) (response.AddToCartRequest, error) {
	var cart entity.Cart
	copier.Copy(&cart, &req)
	err := u.productRepository.AddToCart(cart)
	if err != nil {
		return response.AddToCartRequest{}, nil
	}
	var productResponse response.AddToCartRequest
	copier.Copy(&productResponse, &cart)
	return productResponse, nil
}

func (u ProductUsecase) GetFilteredProduct(category string, hargaterendah string, hargatertinggi string) ([]response.GetProductFilteredResponse, error) {

	products, err := u.productRepository.GetFiltered(category, hargaterendah, hargatertinggi)
	if err != nil {
		return nil, err
	}
	var productResponse []response.GetProductFilteredResponse
	copier.Copy(&productResponse, &products)
	return productResponse, nil
}

func (u ProductUsecase) GetCartList() ([]response.GetCartResponse, error) {
	carts, err := u.productRepository.GetCartAll()
	if err != nil {
		return nil, nil
	}
	var cartsResponse []response.GetCartResponse
	copier.Copy(&cartsResponse, &carts)
	return cartsResponse, nil
}
