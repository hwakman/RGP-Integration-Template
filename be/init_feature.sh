TYPENAME=$1
NAME=$2

cat > ./features/$NAME.go << EOF
package features

import "github.com/labstack/echo/v4"

type ${TYPENAME} struct {}

func (h ${TYPENAME}) ${TYPENAME}Route(c echo.Context) error {
	return c.String(200, "${TYPENAME} route")
}
EOF
