package main

import (
	"efishery-ecommerce/config"
	"efishery-ecommerce/handler"
	"efishery-ecommerce/repository"
	"efishery-ecommerce/routes"
	"efishery-ecommerce/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.AutoMigrate()
	//err := config.Seeding()
	//if err != nil {
	//	panic(err)
	//}

	app := echo.New()

	productRepository := repository.NewProductRepository(config.DB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	routes.Routes(app, productHandler)
	app.Logger.Fatal(app.Start(":8080"))
}
