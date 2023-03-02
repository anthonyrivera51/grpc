package rutas

import "github.com/labstack/echo/v4"

type Rutas struct {
	echoInstance *echo.Group
}

func NewRutas(echoInstance *echo.Group) *Rutas {
	return &Rutas{
		echoInstance: echoInstance,
	}
}
