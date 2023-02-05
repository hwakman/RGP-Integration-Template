package features

import "github.com/labstack/echo/v4"

type Home struct {}

func (h Home) HomeRoute(c echo.Context) error {
	return c.String(200, "Home route")
}
