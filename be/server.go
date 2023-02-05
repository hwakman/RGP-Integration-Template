package main

import (
	"echoapp/features"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello");
}

func main() {
	app := echo.New()

	// home
	home := new(features.Home)
	app.GET("/", home.HomeRoute)

	// product
	product := new(features.Product)
	app.GET("/product", product.ProductRoute)

	// brand
	brand := new(features.Brand)
	app.GET("/brand", brand.BrandRoute)

	// order
	order := new(features.Order)
	app.GET("/order", order.OrderRoute)

	// log
	app.Logger.Fatal(app.Start(":8000"))
}
