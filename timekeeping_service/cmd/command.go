package main

import (
	"Backend-Server/library/server"
	"fmt"
	"github.com/urfave/cli/v2"
)

func serverAction(cliCtx *cli.Context) error {
	service, err := newService(cfg)
	if err != nil {
		return err
	}
	s, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithHttpAddrListen(cfg.Server.HTTP),
		server.WithServiceServer(service),
	)
	if err != nil {
		return err
	}
	if err := s.Serve(); err != nil {
		return fmt.Errorf("Error start servers. %w", err)
	}
	return nil
}

func configAction(context *cli.Context) error {
	if cfg == nil {
		return err
	}
	fmt.Printf("Config: %+v\n", cfg)
	return nil
}
