package handler

import (
	"efishery-ecommerce/entity/response"
	"efishery-ecommerce/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productcase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productcase}
}

func (h ProductHandler) GetProductList(ctx echo.Context) error {
	products, err := h.productUsecase.GetProductList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list product failed",
			Data:    nil,
		})
	}
	if len(products) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no products found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully get list all products",
		Data:    products,
	})

}

func (h ProductHandler) GetProductDetail(ctx echo.Context) error {
	id := ctx.Param("id")
	products, err := h.productUsecase.GetProductDetail(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list product failed",
			Data:    nil,
		})
	}
	if len(products) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no products found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully get list all products",
		Data:    products,
	})
}

func (h ProductHandler) GetProductFilter(ctx echo.Context) error {
	category := ctx.QueryParam("kategori")
	hargaterendah := ctx.QueryParam("hargaterendah")
	hargatertinggi := ctx.QueryParam("hargatertinggi")
	products, err := h.productUsecase.GetFilteredProduct(category, hargaterendah, hargatertinggi)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get filtered product failed",
			Data:    nil,
		})
	}
	if len(products) == 0 {
		return ctx.JSON(http.StatusFound, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no products found",
			Data:    map[string]interface{}{},
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully get all filtered products",
		Data:    products,
	})
}

func (h ProductHandler) AddToCart(ctx echo.Context) error {
	productRequest := new(response.AddToCartRequest)
	if err := ctx.Bind(&productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	fmt.Println(*productRequest)
	productResponse, err := h.productUsecase.AddProductToCart(productRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to add product to cart",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "product added to cart successfully",
		Data:    productResponse,
	})
}

func (h ProductHandler) GetCartList(ctx echo.Context) error {
	carts, err := h.productUsecase.GetCartList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list product failed",
			Data:    nil,
		})
	}
	if len(carts) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no products found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully get list all products",
		Data:    carts,
	})
}

func (h ProductHandler) PostPayment(ctx echo.Context) error {
	username := ctx.FormValue("username")
	//-----------
	// Read file
	//-----------
	// Source
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	filename := "./static/" + username + "_proof.pdf" // ini harus diubah sesuai hosting nanti
	proofRequest := response.PostProofResponse{
		Name:         username,
		Bukti:        filename,
		IsCheckedOut: true,
	}
	err = h.productUsecase.UploadProof(proofRequest, file)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "upload payment proof failed",
			Data:    err,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully uploaded proof",
		Data:    "saved at: " + filename,
	})
}
