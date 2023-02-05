package features

import "github.com/labstack/echo/v4"

type Product struct {}

func (h Product) ProductRoute(c echo.Context) error {
	return c.String(200, "Product route")
}
