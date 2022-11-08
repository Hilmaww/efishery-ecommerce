package routes

import (
	"efishery-ecommerce/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(app *echo.Echo, productHandler *handler.ProductHandler) {
	r := app.Group("v1")

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r.GET("/products", productHandler.GetProductList)
	r.GET("/products/:id", productHandler.GetProductDetail)
	//r.GET("/products", productHandler.GetProductPerCategory)
	r.POST("/cart", productHandler.AddToCart)
	//r.GET("/cart", productHandler.GetCartList)
	//r.GET("/payment", productHandler.PostPayment)

}
