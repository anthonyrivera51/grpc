package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/anthonyrivera51/grpc-hello-word/internal/infrastructure/enums"
	apiRestfactory "github.com/anthonyrivera51/grpc-hello-word/internal/infrastructure/factory/api_rest.factory"
	"github.com/joho/godotenv"
)

func bootstrap() error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var wg sync.WaitGroup
	wg.Add(1)

	apiRestFactory, err := apiRestfactory.APIRestFactory(enums.Echo)

	if err != nil {
		return fmt.Errorf("failed to setup Echo: %v", err)
	}

	err = apiRestFactory.RunServer()

	if err != nil {
		panic("Error running server")
	}

	wg.Wait()

	return nil

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Err: ", err)
	}

	if err := bootstrap(); err != nil {
		panic(err)
	}
}
