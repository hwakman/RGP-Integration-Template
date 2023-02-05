package features

import "github.com/labstack/echo/v4"

type Order struct {}

func (h Order) OrderRoute(c echo.Context) error {
	return c.String(200, "Order route")
}
