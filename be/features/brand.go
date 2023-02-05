package features

import "github.com/labstack/echo/v4"

type Brand struct {}

func (h Brand) BrandRoute(c echo.Context) error {
	return c.String(200, "Brand route")
}
