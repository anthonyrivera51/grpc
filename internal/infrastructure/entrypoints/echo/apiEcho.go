package apiecho

import (
	"os"

	"github.com/anthonyrivera51/grpc-hello-word/internal/infrastructure/entrypoints/echo/rutas"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

type ApiEcho struct {
	echoInstance  *echo.Echo
	routes        *rutas.Rutas
	globalGroupV1 *echo.Group
}

func NewApiEcho() *ApiEcho {
	var e *echo.Echo = echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	var globalV1 *echo.Group = e.Group("/v1")

	return &ApiEcho{
		echoInstance:  e,
		routes:        rutas.NewRutas(globalV1),
		globalGroupV1: globalV1,
	}
}

type UserServiceServer struct{}

func (api *ApiEcho) RunServer() error {
	api.routes.OnPath()

	if os.Getenv("APP_PORT") == "" {
		os.Setenv("APP_PORT", "50051")
	}

	api.echoInstance.Logger.Fatal(api.echoInstance.Start(":" + os.Getenv("APP_PORT")))

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	srv := &UserServiceServer{}

	gen.
	return nil
}
