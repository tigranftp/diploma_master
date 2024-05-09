package main

import (
	"context"
	"log"
	"nn_auto/api"
	"nn_auto/db/postgres"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	storage, err := postgres.New("postgres://postgres:qwerty@localhost:11235/postgres", context.Background())
	if err != nil {
		return err
	}

	api, err := api.NewAPI(storage)
	if err != nil {
		return err
	}
	defer api.Stop()
	return api.Start()
}
