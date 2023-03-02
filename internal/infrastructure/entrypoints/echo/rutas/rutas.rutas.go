package rutas

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (r *Rutas) OnPath() {
	v1 := r.echoInstance.Group("/Users")
	v1.POST("/login", r.Login)
}

func (r *Rutas) Login(c echo.Context) error {
	fmt.Println("Post Login")

	return nil
}
