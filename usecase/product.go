package usecase

import (
	"efishery-ecommerce/entity"
	"efishery-ecommerce/entity/response"
	"efishery-ecommerce/repository"
	"fmt"
	"github.com/jinzhu/copier"
	"io"
	"mime/multipart"
	"os"
)

type IProductUsecase interface {
	GetProductList() ([]response.GetProductResponse, error)
	GetProductDetail(id string) ([]response.GetProductDetailResponse, error)
	GetFilteredProduct(category string, hargaterendah string, hargatertinggi string) ([]response.GetProductResponse, error)
	AddProductToCart(req *response.AddToCartRequest) (response.AddToCartRequest, error)
	GetCartList() ([]response.GetCartResponse, error)
	UploadProof(proofRequest response.PostProofResponse, file *multipart.FileHeader) error
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

func (u ProductUsecase) UploadProof(proofRequest response.PostProofResponse, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename := proofRequest.Bukti // ini harus diubah sesuai hosting nanti
	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println("error in creation")
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	var proof entity.Proof
	copier.Copy(&proof, &proofRequest)
	err = u.productRepository.UploadProofDirectory(proof)
	if err != nil {
		return err
	}
	err = u.productRepository.DeleteCartList()
	if err != nil {
		return err
	}
	return nil
}
