package handler

import (
	"efishery-ecommerce/entity/response"
	"efishery-ecommerce/usecase"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
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

//func (h ProductHandler) GetProductPerCategory(ctx echo.Context) error {
//	category := ctx.QueryParam("category")
//	products, err := h.productUsecase.GetProductDetail(id)
//	if err != nil {
//		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
//			Code:    http.StatusBadRequest,
//			Message: "get list product failed",
//			Data:    nil,
//		})
//	}
//	if len(products) < 0 {
//		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
//			Code:    http.StatusNoContent,
//			Message: "no products found",
//			Data:    nil,
//		})
//	}
//	return ctx.JSON(http.StatusOK, response.BaseResponse{
//		Code:    http.StatusOK,
//		Message: "successfully get list all products",
//		Data:    products,
//	})
//}

func (h ProductHandler) AddToCart(ctx echo.Context) error {
	var data map[string]interface{} = map[string]interface{}{}
	if err := ctx.Bind(&data); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	var productRequest []response.AddToCartRequest
	dataMarshaled, err := json.Marshal(data)
	if err != nil {
		return err
	}
	dataString := string(dataMarshaled)
	copier.Copy(&productRequest, &dataString)
	fmt.Println(productRequest)
	//productResponse, err := h.productUsecase.AddProductToCart(productRequest)
	//if err != nil {
	//	return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
	//		Code:    http.StatusBadRequest,
	//		Message: "Failed to add product to cart",
	//		Data:    nil,
	//	})
	//}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "product added to cart successfully",
		Data:    nil, //productResponse,
	})
}
