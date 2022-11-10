package routes

import (
	"efishery-ecommerce/handler"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(app *echo.Echo, productHandler *handler.ProductHandler) {
	r := app.Group("v1")

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	r.GET("/test", func(c echo.Context) error {
		id := c.QueryParam("id")
		isIt := id == ""
		fmt.Println(isIt)
		return c.String(http.StatusOK, id)
	})

	r.GET("/products", productHandler.GetProductList)
	r.GET("/products/:id", productHandler.GetProductDetail)
	r.GET("/products/filter", productHandler.GetProductFilter)
	r.POST("/cart", productHandler.AddToCart)
	r.GET("/cart", productHandler.GetCartList)
	r.POST("/payment", productHandler.PostPayment)

}
