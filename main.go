package main

import (
	"context"

	"github.com/phaesoo/pigeonhole/configs"
	"github.com/phaesoo/pigeonhole/internal/logging"
	"github.com/phaesoo/pigeonhole/rpc"

	"golang.org/x/sync/errgroup"
)

func main() {
	config := configs.Get()

	var g errgroup.Group

	logger := logging.New(config.App.Name, config.Log)

	g.Go(func() error {
		s := rpc.NewServer(config.App, logger)
		return s.RunServer(context.Background())
	})
}
