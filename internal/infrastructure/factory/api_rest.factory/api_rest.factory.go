package apirestfactory

import (
	"errors"

	apiecho "github.com/anthonyrivera51/grpc-hello-word/internal/infrastructure/entrypoints/echo"
	"github.com/anthonyrivera51/grpc-hello-word/internal/infrastructure/enums"
	"github.com/anthonyrivera51/grpc-hello-word/internal/infrastructure/factory/interfaces"
)

func APIRestFactory(provider enums.APIRestProvider) (interfaces.APIRestInterface, error) {
	switch provider {
	case enums.Echo:
		return apiecho.NewApiEcho(), nil

	default:
		return nil, errors.New("NoAdapterFound")
	}
}
