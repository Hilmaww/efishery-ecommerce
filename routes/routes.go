package routes

import (
	"efishery-ecommerce/config"
	"efishery-ecommerce/entity/response"
	"efishery-ecommerce/handler"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo, productHandler *handler.ProductHandler) {
	r := app.Group("v1")

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! Efishery Ecommerce Server is On! ")
	})
	r.GET("/test", func(c echo.Context) error {
		id := c.QueryParam("id")
		isIt := id == ""
		fmt.Println(isIt)
		return c.String(http.StatusOK, id)
	})

	r.POST("/seed", func(c echo.Context) error {
		err := config.Seeding()
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "successfully done seeding",
			Data:    "done",
		})
	})

	r.GET("/products", productHandler.GetProductList)
	r.GET("/products/:id", productHandler.GetProductDetail)
	r.GET("/products/filter", productHandler.GetProductFilter)
	r.POST("/cart", productHandler.AddToCart)
	r.GET("/cart", productHandler.GetCartList)
	r.POST("/payment", productHandler.PostPayment)

}
